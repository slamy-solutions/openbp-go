package system

import (
	"os"
)

type CacheConfig struct {
	Enabled bool
	URL     string
}
type NatsConfig struct {
	Enabled    bool
	URL        string
	ClientName string
}
type DBConfig struct {
	Enabled bool
	URL     string
}
type OTelConfig struct {
	Enabled bool
	URL     string

	ServiceModule     string
	ServiceName       string
	ServiceVersion    string
	ServiceInstanceID string
}
type VaultConfig struct {
	Enabled bool
	URL     string
}

func getConfigEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewOTelConfig(module string, name string, version string, instanceID string) *OTelConfig {
	return &OTelConfig{
		Enabled:           true,
		URL:               getConfigEnv("SYSTEM_TELEMETRY_EXPORTER_ENDPOINT", "system_telemetry:55680"),
		ServiceModule:     module,
		ServiceName:       name,
		ServiceVersion:    version,
		ServiceInstanceID: instanceID,
	}
}

func (c *OTelConfig) WithURL(url string) *OTelConfig {
	c.URL = url
	return c
}

type StubConfig struct {
	Cache CacheConfig
	Nats  NatsConfig
	Db    DBConfig
	OTel  OTelConfig
	Vault VaultConfig
}

func (s *StubConfig) WithCache(config ...CacheConfig) *StubConfig {
	cfg := &CacheConfig{
		Enabled: true,
		URL:     getConfigEnv("SYSTEM_CACHE_URL", "redis://system_cache"),
	}
	if len(config) > 0 {
		cfg = &config[0]
	}

	s.Cache = *cfg
	return s
}

func (s *StubConfig) WithNats(config ...NatsConfig) *StubConfig {
	cfg := &NatsConfig{
		Enabled:    true,
		URL:        getConfigEnv("SYSTEM_NATS_URL", "nats://system_nats:4222"),
		ClientName: "",
	}
	if len(config) > 0 {
		cfg = &config[0]
	}

	s.Nats = *cfg
	return s
}

func (s *StubConfig) WithDB(config ...DBConfig) *StubConfig {
	cfg := &DBConfig{
		Enabled: true,
		URL:     getConfigEnv("SYSTEM_DB_URL", "mongodb://root:example@system_db/admin"),
	}
	if len(config) > 0 {
		cfg = &config[0]
	}

	s.Db = *cfg
	return s
}

func (s *StubConfig) WithOTel(config *OTelConfig) *StubConfig {
	s.OTel = *config
	return s
}

func (s *StubConfig) WithVault(config ...VaultConfig) *StubConfig {
	cfg := &VaultConfig{
		Enabled: true,
		URL:     getConfigEnv("SYSTEM_VAULT_URL", "system_vault:80"),
	}
	if len(config) > 0 {
		cfg = &config[0]
	}

	s.Vault = *cfg
	return s
}

func NewSystemStubConfig() *StubConfig {
	return &StubConfig{
		Cache: CacheConfig{
			Enabled: false,
			URL:     "",
		},
		Nats: NatsConfig{
			Enabled:    false,
			URL:        "",
			ClientName: "",
		},
		Db: DBConfig{
			Enabled: false,
			URL:     "",
		},
		OTel: OTelConfig{
			Enabled:           false,
			URL:               "",
			ServiceModule:     "",
			ServiceName:       "",
			ServiceVersion:    "",
			ServiceInstanceID: "",
		},
		Vault: VaultConfig{
			Enabled: false,
			URL:     "",
		},
	}
}
