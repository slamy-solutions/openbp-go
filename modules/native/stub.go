package native

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/actor/user"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/auth"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/authentication/password"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/identity"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/policy"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/role"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/token"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/keyvaluestorage"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/namespace"

	"google.golang.org/grpc"
)

type IAMAuthenticationServices struct {
	Password password.IAMAuthenticationPasswordServiceClient
}

type Stub struct {
	Namespace       namespace.NamespaceServiceClient
	KeyValueStorage keyvaluestorage.KeyValueStorageServiceClient

	ActorUser user.ActorUserServiceClient

	IAMAuthentication IAMAuthenticationServices
	IAMPolicy         policy.IAMPolicyServiceClient
	IAMRole           role.IAMRoleServiceClient
	IAMIdentity       identity.IAMIdentityServiceClient
	IAMToken          token.IAMTokenServiceClient
	IAMAuth           auth.IAMAuthServiceClient

	log       *log.Logger
	config    *StubConfig
	mu        sync.Mutex
	connected bool
	dials     []*grpc.ClientConn
}

func NewStub(config *StubConfig) *Stub {
	return &Stub{
		log:       config.logger,
		config:    config,
		mu:        sync.Mutex{},
		connected: false,
		dials:     []*grpc.ClientConn{},
	}
}

func connectToService[T interface{}](s *Stub, clientFunction func(grpc.ClientConnInterface) T, config *GrpcServiceConfig, serviceName string) (T, error) {
	conn, service, err := makeGrpcClient(clientFunction, config.url)
	if err != nil {
		s.closeConnections()
		s.log.Error("Error while connecting to the [" + serviceName + "] service: " + err.Error())
		return service, err
	}
	s.log.Info("Successfully connected to the [" + serviceName + "] service")
	s.dials = append(s.dials, conn)
	return service, err
}

func (s *Stub) Connect() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.connected {
		return nil
	}

	if s.config.namespace.enabled {
		service, err := connectToService(s, namespace.NewNamespaceServiceClient, &s.config.namespace, "native_namespace")
		if err != nil {
			return err
		}
		s.Namespace = service
	}

	if s.config.keyValueStorage.enabled {
		service, err := connectToService(s, keyvaluestorage.NewKeyValueStorageServiceClient, &s.config.keyValueStorage, "native_actor_keyvaluestorage")
		if err != nil {
			return err
		}
		s.KeyValueStorage = service
	}

	if s.config.actorUser.enabled {
		service, err := connectToService(s, user.NewActorUserServiceClient, &s.config.actorUser, "native_actor_user")
		if err != nil {
			return err
		}
		s.ActorUser = service
	}

	if s.config.iamAuthentication.enabled {
		dial, err := makeGrpcDial(s.config.iamAuthentication.url)
		if err != nil {
			s.closeConnections()
			s.log.Error("Error while connecting to the [native_iam_authentication] service: " + err.Error())
			return err
		}
		s.log.Info("Successfully connected to the [native_iam_authentication] service")
		s.dials = append(s.dials, dial)

		s.IAMAuthentication = IAMAuthenticationServices{
			Password: password.NewIAMAuthenticationPasswordServiceClient(dial),
		}
	}

	if s.config.iamIdentity.enabled {
		service, err := connectToService(s, identity.NewIAMIdentityServiceClient, &s.config.iamIdentity, "native_iam_identity")
		if err != nil {
			return err
		}
		s.IAMIdentity = service
	}

	if s.config.iamPolicy.enabled {
		service, err := connectToService(s, policy.NewIAMPolicyServiceClient, &s.config.iamPolicy, "native_iam_policy")
		if err != nil {
			return err
		}
		s.IAMPolicy = service
	}

	if s.config.iamRole.enabled {
		service, err := connectToService(s, role.NewIAMRoleServiceClient, &s.config.iamRole, "native_iam_role")
		if err != nil {
			return err
		}
		s.IAMRole = service
	}

	if s.config.iamAuth.enabled {
		service, err := connectToService(s, auth.NewIAMAuthServiceClient, &s.config.iamAuth, "native_iam_auth")
		if err != nil {
			return err
		}
		s.IAMAuth = service
	}

	if s.config.iamToken.enabled {
		service, err := connectToService(s, token.NewIAMTokenServiceClient, &s.config.iamToken, "native_iam_token")
		if err != nil {
			return err
		}
		s.IAMToken = service
	}

	s.connected = true
	return nil
}

func (s *Stub) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.connected {
		return
	}

	s.closeConnections()
}

func (s *Stub) closeConnections() {
	for _, dial := range s.dials {
		dial.Close()
	}
	s.dials = make([]*grpc.ClientConn, 0)
	s.connected = false
}
