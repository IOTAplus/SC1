// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type InitCall struct {
	Func   *wasmlib.ScInitFunc
	Params MutableInitParams
}

type MintCall struct {
	Func   *wasmlib.ScFunc
	Params MutableMintParams
}

type TransferCall struct {
	Func   *wasmlib.ScFunc
	Params MutableTransferParams
}

type BalanceOfCall struct {
	Func    *wasmlib.ScView
	Params  MutableBalanceOfParams
	Results ImmutableBalanceOfResults
}

type TotalSupplyCall struct {
	Func    *wasmlib.ScView
	Results ImmutableTotalSupplyResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit, keyMap[:], idxMap[:])}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Mint(ctx wasmlib.ScFuncCallContext) *MintCall {
	f := &MintCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncMint)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Transfer(ctx wasmlib.ScFuncCallContext) *TransferCall {
	f := &TransferCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncTransfer)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) BalanceOf(ctx wasmlib.ScViewCallContext) *BalanceOfCall {
	f := &BalanceOfCall{Func: wasmlib.NewScView(ctx, HScName, HViewBalanceOf)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) TotalSupply(ctx wasmlib.ScViewCallContext) *TotalSupplyCall {
	f := &TotalSupplyCall{Func: wasmlib.NewScView(ctx, HScName, HViewTotalSupply)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}
