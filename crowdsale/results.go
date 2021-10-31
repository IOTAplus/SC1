// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package crowdsale

 import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableGetOwnerResults struct {
	id int32
}

func (s ImmutableGetOwnerResults) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxResultOwner])
}

type MutableGetOwnerResults struct {
	id int32
}

func (s MutableGetOwnerResults) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxResultOwner])
}
