package openbp

import (
	"context"

	"github.com/slamy-solutions/openbp-go/modules"
)

func ConnectToModules(ctx context.Context, config *modules.StubConfig) (*modules.OpenBPStub, error) {
	stub := modules.NewOpenBPStub(config)
	return stub, stub.Connect(ctx)
}
