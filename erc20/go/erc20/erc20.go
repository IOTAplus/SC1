// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package erc20

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"
)

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	supply := f.Params.Supply().Value()
	ctx.Require(supply > 0, "ERC20: supply must be greater than 0")
	f.State.Supply().SetValue(supply)

	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
	} else {
		f.State.Owner().SetValue(ctx.ContractCreator())
	}

	owner := f.State.Owner().Value()
	f.State.Balances().GetInt64(owner).SetValue(supply)
}

func funcMint(ctx wasmlib.ScFuncContext, f *MintContext) {
	amount := f.Params.Amount().Value()
	ctx.Require(amount > 0, "erc20.mint.fail: wrong 'amount' parameter")

	recipient := f.Params.To().Value()

	balances := f.State.Balances()
	recipientBalance := balances.GetInt64(recipient)

	recipientBalance.SetValue(recipientBalance.Value() + amount)
	f.State.Supply().SetValue(f.State.Supply().Value() + amount)
}

func funcTransfer(ctx wasmlib.ScFuncContext, f *TransferContext) {
	amount := f.Params.Amount().Value()
	ctx.Require(amount > 0, "erc20.transfer.fail: wrong 'amount' parameter")

	balances := f.State.Balances()
	sourceBalance := balances.GetInt64(ctx.Caller())
	ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer.fail: not enough funds")

	targetAddr := f.Params.Account().Value()
	targetBalance := balances.GetInt64(targetAddr)
	result := targetBalance.Value() + amount
	ctx.Require(result > 0, "erc20.transfer.fail: overflow")

	sourceBalance.SetValue(sourceBalance.Value() - amount)
	targetBalance.SetValue(targetBalance.Value() + amount)

	// ctx.Event(fmt.Sprintf("Tranfer %d IEXP tokens to account %d", amount, targetBalance.Value()))
}

func viewBalanceOf(ctx wasmlib.ScViewContext, f *BalanceOfContext) {
	account := f.Params.Account().Value()
	balances := f.State.Balances()
	balance := balances.GetInt64(account)
	f.Results.Amount().SetValue(balance.Value())
}

func viewTotalSupply(ctx wasmlib.ScViewContext, f *TotalSupplyContext) {
	f.Results.Supply().SetValue(f.State.Supply().Value())
}
