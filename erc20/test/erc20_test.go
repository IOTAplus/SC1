package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaplus/SC1/erc20/go/erc20"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {

	chain := wasmsolo.StartChain(t, "chain1")
	creator := wasmsolo.NewSoloAgent(chain.Env)

	init := erc20.ScFuncs.Init(nil)
	init.Params.Supply().SetValue(initialSupply)
	init.Params.Owner().SetValue(creator.ScAgentID())

	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)

	require.NoError(t, ctx.ContractExists(erc20.ScName))
}
