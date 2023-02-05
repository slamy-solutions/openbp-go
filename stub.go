package openbp

import "github.com/slamy-solutions/openbp-go/modules"

func ConnectToModules(config *modules.StubConfig) (*modules.OpenBPStub, error) {
	stub := modules.NewOpenBPStub(config)
	return stub, stub.Connect()
}
