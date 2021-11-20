package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaplus/SC1/erc20/go/erc20"
	"github.com/stretchr/testify/require"
)

func getChainAndCreator(t *testing.T) (*solo.Chain, *wasmsolo.SoloAgent) {
	chain := wasmsolo.StartChain(t, "chain1")
	creator := wasmsolo.NewSoloAgent(chain.Env)
	return chain, creator
}

func buildInit(creator *wasmsolo.SoloAgent) *erc20.InitCall {
	init := erc20.ScFuncs.Init(nil)
	init.Params.Supply().SetValue(initialSupply)
	init.Params.Owner().SetValue(creator.ScAgentID())
	return init
}

func setupInit(t *testing.T) *erc20.InitCall {
	_, creator := getChainAndCreator(t)
	init := buildInit(creator)
	return init
}

func TestDeploy(t *testing.T) {
	init := setupInit(t)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)
	require.NoError(t, ctx.ContractExists(erc20.ScName))
}

func TestTotalSupply(t *testing.T) {
	init := setupInit(t)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)

	totalSupply := erc20.ScFuncs.TotalSupply(ctx)
	totalSupply.Func.Call()
	require.NoError(t, ctx.Err)

	supply := totalSupply.Results.Supply()
	require.True(t, supply.Exists())
	require.EqualValues(t, initialSupply, supply.Value())
}
