package native

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func getConfigEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type GrpcServiceConfig struct {
	enabled bool
	url     string
}

func NewGrpcServiceConfig() GrpcServiceConfig {
	return GrpcServiceConfig{
		enabled: false,
		url:     "",
	}
}

func (c *GrpcServiceConfig) Enabled(enabled bool) *GrpcServiceConfig {
	c.enabled = enabled
	return c
}

func (c *GrpcServiceConfig) WithURL(url string) *GrpcServiceConfig {
	c.url = url
	return c
}

type StubConfig struct {
	logger *log.Logger

	namespace       GrpcServiceConfig
	keyValueStorage GrpcServiceConfig
	actorUser       GrpcServiceConfig

	iamAuthenticationPassword GrpcServiceConfig
	iamIdentity               GrpcServiceConfig
	iamAuth                   GrpcServiceConfig
	iamPolicy                 GrpcServiceConfig
	iamRole                   GrpcServiceConfig
	iamToken                  GrpcServiceConfig
}

func NewstubConfig() *StubConfig {
	return &StubConfig{
		logger:                    log.StandardLogger(),
		namespace:                 NewGrpcServiceConfig(),
		keyValueStorage:           NewGrpcServiceConfig(),
		actorUser:                 NewGrpcServiceConfig(),
		iamAuthenticationPassword: NewGrpcServiceConfig(),
		iamIdentity:               NewGrpcServiceConfig(),
		iamAuth:                   NewGrpcServiceConfig(),
		iamPolicy:                 NewGrpcServiceConfig(),
		iamRole:                   NewGrpcServiceConfig(),
		iamToken:                  NewGrpcServiceConfig(),
	}
}

func (sc *StubConfig) WithLogger(logger *log.Logger) *StubConfig {
	sc.logger = logger
	return sc
}

func (sc *StubConfig) WithActorUserService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.actorUser = conf[0]
	} else {
		sc.actorUser = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_ACTOR_USER_URL", "native_actor_user:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithNamespaceService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.namespace = conf[0]
	} else {
		sc.namespace = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_NAMESPACE_URL", "native_namespace:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithKeyValueStorageService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.keyValueStorage = conf[0]
	} else {
		sc.keyValueStorage = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_KEYVALUESTORAGE_URL", "native_keyvaluestorage:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithIAMAuthenticationPasswordService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iamAuthenticationPassword = conf[0]
	} else {
		sc.iamAuthenticationPassword = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_AUTHENTICATION_PASSWORD_URL", "native_iam_authentication_password:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithIAMIdentityService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iamIdentity = conf[0]
	} else {
		sc.iamIdentity = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_IDENTITY_URL", "native_iam_identity:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithIAMPolicyService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iamPolicy = conf[0]
	} else {
		sc.iamPolicy = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_POLICY_URL", "native_iam_policy:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithIAMRoleService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iamRole = conf[0]
	} else {
		sc.iamRole = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_ROLE_URL", "native_iam_role:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithIAMTokenService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iamToken = conf[0]
	} else {
		sc.iamToken = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_TOKEN_URL", "native_iam_token:80"),
		}
	}
	return sc
}

func (sc *StubConfig) WithIAMAuthService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iamAuth = conf[0]
	} else {
		sc.iamAuth = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_AUTH_URL", "native_iam_auth:80"),
		}
	}
	return sc
}
