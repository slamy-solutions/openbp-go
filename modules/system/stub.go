package system

import (
	"context"
	"errors"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/slamy-solutions/openbp-go/modules/system/cache"
	"github.com/slamy-solutions/openbp-go/modules/system/db"
	systemNats "github.com/slamy-solutions/openbp-go/modules/system/nats"
	"github.com/slamy-solutions/openbp-go/modules/system/otel"
	"github.com/slamy-solutions/openbp-go/modules/system/proto/vault"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Stub struct {
	Cache cache.Cache
	DB    *mongo.Client
	OTel  otel.Telemetry
	Nats  *nats.Conn
	Vault vault.VaultServiceClient

	config    *StubConfig
	mu        sync.Mutex
	connected bool
	dials     []*grpc.ClientConn
}

func NewStub(config *StubConfig) *Stub {
	return &Stub{
		config:    config,
		mu:        sync.Mutex{},
		connected: false,
	}
}

func (s *Stub) Connect(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.connected {
		return nil
	}

	if s.config.Vault.Enabled {
		conn, service, err := makeGrpcClient(vault.NewVaultServiceClient, s.config.Vault.URL)
		if err != nil {
			return errors.New("failed to initialize connection to the vault: " + err.Error())
		}
		s.dials = append(s.dials, conn)
		s.Vault = service
	}

	if s.config.OTel.Enabled {
		tel, err := otel.Register(ctx, s.config.OTel.URL, s.config.OTel.ServiceModule, s.config.OTel.ServiceName, s.config.OTel.ServiceVersion, s.config.OTel.ServiceInstanceID)
		if err != nil {
			s.closeGRPCConnections()
			return errors.New("failed to initialize connection to the otel: " + err.Error())
		}
		s.OTel = tel
	}

	if s.config.Cache.Enabled {
		cacheClient, err := cache.New(s.config.Cache.URL)
		if err != nil {
			//Close opened connections
			s.closeGRPCConnections()
			if s.config.OTel.Enabled {
				s.OTel.Shutdown(ctx)
			}

			return errors.New("failed to initialize connection to the cache: " + err.Error())
		}
		s.Cache = cacheClient
	}

	if s.config.Db.Enabled {
		dbClient, err := db.Connect(s.config.Db.URL)
		if err != nil {
			//Close opened connections
			s.closeGRPCConnections()
			if s.config.Cache.Enabled {
				s.Cache.Shutdown(ctx)
			}
			if s.config.OTel.Enabled {
				s.OTel.Shutdown(ctx)
			}

			return errors.New("failed to initialize connection to the DB: " + err.Error())
		}

		s.DB = dbClient
	}

	if s.config.Nats.Enabled {
		natsClient, err := systemNats.Connect(s.config.Nats.URL, s.config.Nats.ClientName)
		if err != nil {
			//Close opened connections
			s.closeGRPCConnections()
			if s.config.Db.Enabled {
				s.DB.Disconnect(ctx)
			}
			if s.config.Cache.Enabled {
				s.Cache.Shutdown(ctx)
			}
			if s.config.OTel.Enabled {
				s.OTel.Shutdown(ctx)
			}

			return errors.New("failed to initialize connection to the Nats: " + err.Error())
		}

		s.Nats = natsClient
	}

	s.connected = true
	return nil
}

func (s *Stub) Close(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.connected {
		return
	}

	if s.config.Cache.Enabled {
		s.Cache.Shutdown(ctx)
	}
	if s.config.Db.Enabled {
		s.DB.Disconnect(ctx)
	}
	if s.config.Nats.Enabled {
		s.Nats.Close()
	}
	if s.config.OTel.Enabled {
		s.OTel.Shutdown(ctx)
	}

	s.closeGRPCConnections()

	s.connected = false
}

func (s *Stub) closeGRPCConnections() {
	for _, dial := range s.dials {
		dial.Close()
	}
	s.dials = make([]*grpc.ClientConn, 0)
}
