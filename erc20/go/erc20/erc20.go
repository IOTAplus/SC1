// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func funcApprove(ctx wasmlib.ScFuncContext, f *ApproveContext) {
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	supply := f.Params.Supply().Value()
	ctx.Require(supply > 0, "ERC20: supply must be greater than 0")
	f.State.Supply().SetValue(supply)

	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
		return
	}
	f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcMint(ctx wasmlib.ScFuncContext, f *MintContext) {
}

func funcTransfer(ctx wasmlib.ScFuncContext, f *TransferContext) {
}

func funcTransferFrom(ctx wasmlib.ScFuncContext, f *TransferFromContext) {
}

func viewAllowance(ctx wasmlib.ScViewContext, f *AllowanceContext) {
}

func viewBalanceOf(ctx wasmlib.ScViewContext, f *BalanceOfContext) {
}

func viewTotalSupply(ctx wasmlib.ScViewContext, f *TotalSupplyContext) {
	f.Results.Supply().SetValue(f.State.Supply().Value())
}
