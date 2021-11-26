// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableERC20State struct {
	id int32
}

func (s ImmutableERC20State) Balances() MapAgentIDToImmutableInt64 {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateBalances], wasmlib.TYPE_MAP)
	return MapAgentIDToImmutableInt64{objID: mapID}
}

func (s ImmutableERC20State) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxStateOwner])
}

func (s ImmutableERC20State) Supply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxStateSupply])
}

type MutableERC20State struct {
	id int32
}

func (s MutableERC20State) Balances() MapAgentIDToMutableInt64 {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateBalances], wasmlib.TYPE_MAP)
	return MapAgentIDToMutableInt64{objID: mapID}
}

func (s MutableERC20State) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxStateOwner])
}

func (s MutableERC20State) Supply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxStateSupply])
}
