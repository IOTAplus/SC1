package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaplus/SC1/erc20/go/erc20"
	"github.com/stretchr/testify/require"
)

func getChainAndAgents(t *testing.T) (*solo.Chain, *wasmsolo.SoloAgent, *wasmsolo.SoloAgent) {
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

func setupInit(t *testing.T) (*wasmsolo.SoloContext, *erc20.InitCall, *wasmsolo.SoloAgent, *wasmsolo.SoloAgent) {
	_, creator, recipient := getChainAndAgents(t)
	init := buildInit(creator)
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad, init.Func)
	return ctx, init, creator, recipient
}

func TestDeploy(t *testing.T) {
	ctx, _, _, _ := setupInit(t)
	require.NoError(t, ctx.ContractExists(erc20.ScName))
}

func TestTotalSupply(t *testing.T) {
	ctx, _, _, _ := setupInit(t)

	totalSupply := erc20.ScFuncs.TotalSupply(ctx)
	totalSupply.Func.Call()
	require.NoError(t, ctx.Err)

	supply := totalSupply.Results.Supply()
	require.True(t, supply.Exists())
	require.EqualValues(t, initialSupply, supply.Value())
}

func TestBalanceOf(t *testing.T) {
	ctx, _, creator, _ := setupInit(t)

	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(creator.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balance := balanceOf.Results.Amount()
	require.True(t, balance.Exists())
	require.EqualValues(t, int64(initialSupply), balance.Value())
}

func TestTransfer(t *testing.T) {
	ctx, _, creator, recipient := setupInit(t)

	amountToTransfer := uint64(100)

	transfer := erc20.ScFuncs.Transfer(ctx)
	transfer.Params.Account().SetValue(recipient.ScAgentID())
	transfer.Params.Amount().SetValue(int64(amountToTransfer))
	transfer.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(creator.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfCreator := balanceOf.Results.Amount()
	require.True(t, balanceOfCreator.Exists())
	require.EqualValues(t, int64(initialSupply)-int64(amountToTransfer), balanceOfCreator.Value())

	// check if recipient got the amount
	balanceOf.Params.Account().SetValue(recipient.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfRecipient := balanceOf.Results.Amount()
	require.True(t, balanceOfRecipient.Exists())
	require.EqualValues(t, int64(amountToTransfer), balanceOfRecipient.Value())
}
