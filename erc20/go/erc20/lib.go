// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncApprove, funcApproveThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncMint, funcMintThunk)
	exports.AddFunc(FuncTransfer, funcTransferThunk)
	exports.AddFunc(FuncTransferFrom, funcTransferFromThunk)
	exports.AddView(ViewAllowance, viewAllowanceThunk)
	exports.AddView(ViewBalanceOf, viewBalanceOfThunk)
	exports.AddView(ViewTotalSupply, viewTotalSupplyThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type ApproveContext struct {
	Params ImmutableApproveParams
	State  MutableERC20State
}

func funcApproveThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcApprove")
	f := &ApproveContext{
		Params: ImmutableApproveParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	ctx.Require(f.Params.Delegation().Exists(), "missing mandatory delegation")
	funcApprove(ctx, f)
	ctx.Log("erc20.funcApprove ok")
}

type InitContext struct {
	Params ImmutableInitParams
	State  MutableERC20State
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Owner().Exists(), "missing mandatory owner")
	ctx.Require(f.Params.Supply().Exists(), "missing mandatory supply")
	funcInit(ctx, f)
	ctx.Log("erc20.funcInit ok")
}

type MintContext struct {
	Params ImmutableMintParams
	State  MutableERC20State
}

func funcMintThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcMint")
	access := ctx.State().GetAgentID(wasmlib.Key("owner"))
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	f := &MintContext{
		Params: ImmutableMintParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	ctx.Require(f.Params.To().Exists(), "missing mandatory to")
	funcMint(ctx, f)
	ctx.Log("erc20.funcMint ok")
}

type TransferContext struct {
	Params ImmutableTransferParams
	State  MutableERC20State
}

func funcTransferThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcTransfer")
	f := &TransferContext{
		Params: ImmutableTransferParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	funcTransfer(ctx, f)
	ctx.Log("erc20.funcTransfer ok")
}

type TransferFromContext struct {
	Params ImmutableTransferFromParams
	State  MutableERC20State
}

func funcTransferFromThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcTransferFrom")
	f := &TransferFromContext{
		Params: ImmutableTransferFromParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	ctx.Require(f.Params.Recipient().Exists(), "missing mandatory recipient")
	funcTransferFrom(ctx, f)
	ctx.Log("erc20.funcTransferFrom ok")
}

type AllowanceContext struct {
	Params  ImmutableAllowanceParams
	Results MutableAllowanceResults
	State   ImmutableERC20State
}

func viewAllowanceThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewAllowance")
	f := &AllowanceContext{
		Params: ImmutableAllowanceParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableAllowanceResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	ctx.Require(f.Params.Delegation().Exists(), "missing mandatory delegation")
	viewAllowance(ctx, f)
	ctx.Log("erc20.viewAllowance ok")
}

type BalanceOfContext struct {
	Params  ImmutableBalanceOfParams
	Results MutableBalanceOfResults
	State   ImmutableERC20State
}

func viewBalanceOfThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewBalanceOf")
	f := &BalanceOfContext{
		Params: ImmutableBalanceOfParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableBalanceOfResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	viewBalanceOf(ctx, f)
	ctx.Log("erc20.viewBalanceOf ok")
}

type TotalSupplyContext struct {
	Results MutableTotalSupplyResults
	State   ImmutableERC20State
}

func viewTotalSupplyThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewTotalSupply")
	f := &TotalSupplyContext{
		Results: MutableTotalSupplyResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableERC20State{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewTotalSupply(ctx, f)
	ctx.Log("erc20.viewTotalSupply ok")
}
