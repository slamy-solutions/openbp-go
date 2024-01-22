package native

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/actor/user"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/auth"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/authentication/oauth2"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/authentication/password"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/authentication/x509"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/identity"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/policy"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/role"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/iam/token"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/keyvaluestorage"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/namespace"

	"github.com/slamy-solutions/openbp-go/modules/native/proto/storage/bucket"
	"github.com/slamy-solutions/openbp-go/modules/native/proto/storage/fs"

	"google.golang.org/grpc"
)

type IamActorServices struct {
	User user.ActorUserServiceClient
}

type IamAuthenticationServices struct {
	Password password.IAMAuthenticationPasswordServiceClient
	X509     x509.IAMAuthenticationX509ServiceClient
	OAuth    IamAuthenticationOAuthServices
}

type IamAuthenticationOAuthServices struct {
	Config oauth2.IAMAuthenticationOAuth2ConfigServiceClient
	OAuth2 oauth2.IAMAuthenticationOAuth2ServiceClient
}

type IAMService struct {
	Actor          *IamActorServices
	Authentication *IamAuthenticationServices
	Identity       identity.IAMIdentityServiceClient
	Auth           auth.IAMAuthServiceClient
	Policy         policy.IAMPolicyServiceClient
	Role           role.IAMRoleServiceClient
	Token          token.IAMTokenServiceClient
}

type StorageService struct {
	Bucket bucket.BucketServiceClient
	FS     fs.FSServiceClient
}

type Stub struct {
	Namespace       namespace.NamespaceServiceClient
	KeyValueStorage keyvaluestorage.KeyValueStorageServiceClient
	IAM             IAMService
	Storage         StorageService

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

	if s.config.iam.enabled {
		dial, err := makeGrpcDial(s.config.iam.url)
		if err != nil {
			s.closeConnections()
			s.log.Error("Error while connecting to the [native_iam] service: " + err.Error())
			return err
		}
		s.log.Info("Successfully connected to the [native_iam] service")
		s.dials = append(s.dials, dial)

		s.IAM = IAMService{
			Actor: &IamActorServices{
				User: user.NewActorUserServiceClient(dial),
			},
			Authentication: &IamAuthenticationServices{
				Password: password.NewIAMAuthenticationPasswordServiceClient(dial),
				X509:     x509.NewIAMAuthenticationX509ServiceClient(dial),
				OAuth: IamAuthenticationOAuthServices{
					Config: oauth2.NewIAMAuthenticationOAuth2ConfigServiceClient(dial),
					OAuth2: oauth2.NewIAMAuthenticationOAuth2ServiceClient(dial),
				},
			},
			Identity: identity.NewIAMIdentityServiceClient(dial),
			Auth:     auth.NewIAMAuthServiceClient(dial),
			Policy:   policy.NewIAMPolicyServiceClient(dial),
			Role:     role.NewIAMRoleServiceClient(dial),
			Token:    token.NewIAMTokenServiceClient(dial),
		}
	}

	if s.config.storage.enabled {
		dial, err := makeGrpcDial(s.config.storage.url)
		if err != nil {
			s.closeConnections()
			s.log.Error("Error while connecting to the [native_storage] service: " + err.Error())
			return err
		}
		s.log.Info("Successfully connected to the [native_storage] service")
		s.dials = append(s.dials, dial)

		s.Storage = StorageService{
			Bucket: bucket.NewBucketServiceClient(dial),
			FS:     fs.NewFSServiceClient(dial),
		}
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
