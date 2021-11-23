package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaplus/SC1/erc20/go/erc20"
	"github.com/stretchr/testify/require"
)

func getChainAndCreator(t *testing.T) (*solo.Chain, *wasmsolo.SoloAgent, *wasmsolo.SoloAgent) {
	chain := wasmsolo.StartChain(t, "chain1")
	creator := wasmsolo.NewSoloAgent(chain.Env)
	recipient := wasmsolo.NewSoloAgent(chain.Env)
	return chain, creator, recipient
}

func buildInit(creator *wasmsolo.SoloAgent) *erc20.InitCall {
	init := erc20.ScFuncs.Init(nil)
	init.Params.Supply().SetValue(initialSupply)
	init.Params.Owner().SetValue(creator.ScAgentID())
	return init
}

func setupInit(t *testing.T) (*erc20.InitCall, *wasmsolo.SoloAgent, *wasmsolo.SoloAgent) {
	_, creator, recipient := getChainAndCreator(t)
	init := buildInit(creator)
	return init, creator, recipient
}

func TestDeploy(t *testing.T) {
	init, _, _ := setupInit(t)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)
	require.NoError(t, ctx.ContractExists(erc20.ScName))
}

func TestTotalSupply(t *testing.T) {
	init, _, _ := setupInit(t)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)

	totalSupply := erc20.ScFuncs.TotalSupply(ctx)
	totalSupply.Func.Call()
	require.NoError(t, ctx.Err)

	supply := totalSupply.Results.Supply()
	require.True(t, supply.Exists())
	require.EqualValues(t, initialSupply, supply.Value())
}

func TestBalanceOf(t *testing.T) {
	init, creator, _ := setupInit(t)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)

	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(creator.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balance := balanceOf.Results.Amount()
	require.True(t, balance.Exists())
	require.EqualValues(t, initialSupply, balance.Value())
}

func TestTransfer(t *testing.T) {
	init, creator, recipient := setupInit(t)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)

	amountToTransfer := int64(100)

	transfer := erc20.ScFuncs.Transfer(ctx)
	transfer.Params.Account().SetValue(recipient.ScAgentID())
	transfer.Params.Amount().SetValue(amountToTransfer)
	transfer.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(creator.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfCreator := balanceOf.Results.Amount()
	require.True(t, balanceOfCreator.Exists())
	require.EqualValues(t, initialSupply-amountToTransfer, balanceOfCreator.Value())

	// check if recipient got the amount
	balanceOf.Params.Account().SetValue(recipient.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfRecipient := balanceOf.Results.Amount()
	require.True(t, balanceOfRecipient.Exists())
	require.EqualValues(t, amountToTransfer, balanceOfRecipient.Value())
}
