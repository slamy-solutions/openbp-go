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
	iam             GrpcServiceConfig
	storage         GrpcServiceConfig
}

func NewstubConfig() *StubConfig {
	return &StubConfig{
		logger:          log.StandardLogger(),
		namespace:       NewGrpcServiceConfig(),
		keyValueStorage: NewGrpcServiceConfig(),
		iam:             NewGrpcServiceConfig(),
		storage:         NewGrpcServiceConfig(),
	}
}

func (sc *StubConfig) WithLogger(logger *log.Logger) *StubConfig {
	sc.logger = logger
	return sc
}

func (sc *StubConfig) WithIAMService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.iam = conf[0]
	} else {
		sc.iam = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_IAM_URL", "native_iam:80"),
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

func (sc *StubConfig) WithStorageService(conf ...GrpcServiceConfig) *StubConfig {
	if len(conf) != 0 {
		sc.storage = conf[0]
	} else {
		sc.storage = GrpcServiceConfig{
			enabled: true,
			url:     getConfigEnv("NATIVE_STORAGE_URL", "native_storage:80"),
		}
	}
	return sc
}
