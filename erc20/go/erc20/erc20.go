// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func funcApprove(ctx wasmlib.ScFuncContext, f *ApproveContext) {
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
		return
	}
	f.State.Owner().SetValue(ctx.ContractCreator())
	f.State.Supply().SetValue(f.Params.Supply().Value())
	f.State.Name().SetValue(f.Params.Name().Value())
	f.State.Symbol().SetValue(f.Params.Symbol().Value())
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

func viewDecimals(ctx wasmlib.ScViewContext, f *DecimalsContext) {
}

func viewName(ctx wasmlib.ScViewContext, f *NameContext) {
	f.Results.Name().SetValue(f.State.Name().Value())
}

func viewSymbol(ctx wasmlib.ScViewContext, f *SymbolContext) {
	f.Results.Symbol().SetValue(f.State.Symbol().Value())
}

func viewTotalSupply(ctx wasmlib.ScViewContext, f *TotalSupplyContext) {
	f.Results.Supply().SetValue(f.State.Supply().Value())
}
