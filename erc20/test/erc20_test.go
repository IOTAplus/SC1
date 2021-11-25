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

func setupInit(t *testing.T) (*wasmsolo.SoloContext, *erc20.InitCall, *wasmsolo.SoloAgent) {
	chain, creator := getChainAndCreator(t)
	init := buildInit(creator)
	ctx := wasmsolo.NewSoloContextForChain(t, chain, nil, erc20.ScName, erc20.OnLoad, init.Func)
	return ctx, init, creator
}

func TestDeploy(t *testing.T) {
	ctx, _, _ := setupInit(t)
	require.NoError(t, ctx.ContractExists(erc20.ScName))
}

func TestTotalSupply(t *testing.T) {
	ctx, _, _ := setupInit(t)

	totalSupply := erc20.ScFuncs.TotalSupply(ctx)
	totalSupply.Func.Call()
	require.NoError(t, ctx.Err)

	supply := totalSupply.Results.Supply()
	require.True(t, supply.Exists())
	require.EqualValues(t, initialSupply, supply.Value())
}

func TestBalanceOf(t *testing.T) {
	ctx, _, creator := setupInit(t)

	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(creator.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balance := balanceOf.Results.Amount()
	require.True(t, balance.Exists())
	require.EqualValues(t, int64(initialSupply), balance.Value())
}

func TestTransfer(t *testing.T) {
	ctx, _, creator := setupInit(t)
	recipient1 := ctx.NewSoloAgent()
	recipient2 := ctx.NewSoloAgent()

	amountToTransfer1 := uint64(100)

	transfer := erc20.ScFuncs.Transfer(ctx.Sign(creator))
	transfer.Params.Account().SetValue(recipient1.ScAgentID())
	transfer.Params.Amount().SetValue(int64(amountToTransfer1))
	transfer.Func.Post()
	transfer.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(creator.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfCreator := balanceOf.Results.Amount()
	require.True(t, balanceOfCreator.Exists())
	require.EqualValues(t, int64(initialSupply)-int64(amountToTransfer1), balanceOfCreator.Value())

	// check if recipient got the amount
	balanceOf.Params.Account().SetValue(recipient1.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfRecipient1 := balanceOf.Results.Amount()
	require.True(t, balanceOfRecipient1.Exists())
	require.EqualValues(t, int64(amountToTransfer1), balanceOfRecipient1.Value())

	// transfer tokens from recipient1 to recipient2
	amountToTransfer2 := uint64(30)

	transfer = erc20.ScFuncs.Transfer(ctx.Sign(recipient1))
	transfer.Params.Account().SetValue(recipient2.ScAgentID())
	transfer.Params.Amount().SetValue(int64(amountToTransfer2))
	transfer.Func.Post()
	transfer.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	balanceOf.Params.Account().SetValue(recipient2.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfRecipient2 := balanceOf.Results.Amount()
	require.True(t, balanceOfRecipient2.Exists())
	require.EqualValues(t, int64(amountToTransfer2), balanceOfRecipient2.Value())

	balanceOf.Params.Account().SetValue(recipient1.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)

	balanceOfRecipient1 = balanceOf.Results.Amount()
	require.True(t, balanceOfRecipient1.Exists())
	require.EqualValues(t, int64(amountToTransfer1)-int64(amountToTransfer2), balanceOfRecipient1.Value())
}

func TestTransferNotEnoughFunds(t *testing.T) {
	ctx, _, creator := setupInit(t)
	recipient1 := ctx.NewSoloAgent()

	amountToTransfer1 := uint64(200)

	transfer := erc20.ScFuncs.Transfer(ctx.Sign(creator))
	transfer.Params.Account().SetValue(recipient1.ScAgentID())
	transfer.Params.Amount().SetValue(int64(amountToTransfer1))
	transfer.Func.Post()
	transfer.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	// transfer tokens from recipient1 to creator
	amountToTransfer2 := uint64(210)

	transfer = erc20.ScFuncs.Transfer(ctx.Sign(recipient1))
	transfer.Params.Account().SetValue(creator.ScAgentID())
	transfer.Params.Amount().SetValue(int64(amountToTransfer2))
	transfer.Func.Post()
	transfer.Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
}
