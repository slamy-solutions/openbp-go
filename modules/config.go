package modules

import (
	log "github.com/sirupsen/logrus"

	"github.com/slamy-solutions/openbp-go/modules/native"
	"github.com/slamy-solutions/openbp-go/modules/system"
)

type StubConfig struct {
	logger *log.Logger
	native *native.StubConfig
	system *system.StubConfig
}

func NewStubConfig() *StubConfig {
	return &StubConfig{
		logger: log.StandardLogger(),
		native: nil,
		system: nil,
	}
}

func (sc *StubConfig) WithLogger(logger *log.Logger) *StubConfig {
	sc.logger = logger
	return sc
}

func (sc *StubConfig) WithNativeModule(config *native.StubConfig) *StubConfig {
	sc.native = config
	return sc
}

func (sc *StubConfig) WithSystemModule(config *system.StubConfig) *StubConfig {
	sc.system = config
	return sc
}
