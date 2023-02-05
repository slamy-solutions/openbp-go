package modules

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/slamy-solutions/openbp-go/modules/native"
	"github.com/slamy-solutions/openbp-go/modules/system"
)

type OpenBPStub struct {
	Native *native.Stub
	System *system.Stub

	log       *log.Logger
	config    *StubConfig
	mu        sync.Mutex
	connected bool
}

func NewOpenBPStub(config *StubConfig) *OpenBPStub {
	return &OpenBPStub{
		log:       config.logger,
		config:    config,
		mu:        sync.Mutex{},
		connected: false,
		Native:    nil,
	}
}

func (s *OpenBPStub) Connect(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.connected {
		return nil
	}

	if s.config.system != nil {
		s.System = system.NewStub(s.config.system)
		if err := s.System.Connect(ctx); err != nil {
			s.log.Error("Failed to connect to the system module. " + err.Error())
			s.closeConnections(context.Background())
			return err
		}
		s.log.Info("Connected to the system module")
	}

	if s.config.native != nil {
		s.Native = native.NewStub(s.config.native)
		if err := s.Native.Connect(); err != nil {
			s.log.Error("Failed to connect to the native module. " + err.Error())
			s.closeConnections(context.Background())
			return err
		}
		s.log.Info("Connected to the native module")
	}

	return nil
}

func (s *OpenBPStub) Close(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.connected {
		return
	}

	s.closeConnections(ctx)
}

func (s *OpenBPStub) closeConnections(ctx context.Context) {
	if s.Native != nil {
		s.Native.Close()
	}
	if s.System != nil {
		s.System.Close(ctx)
	}
	s.connected = false
}
