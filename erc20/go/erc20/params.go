// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20
import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableApproveParams struct {
	id int32
}

func (s ImmutableApproveParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s ImmutableApproveParams) Delegation() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamDelegation])
}

type MutableApproveParams struct {
	id int32
}

func (s MutableApproveParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s MutableApproveParams) Delegation() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamDelegation])
}

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamOwner])
}

func (s ImmutableInitParams) Supply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamSupply])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamOwner])
}

func (s MutableInitParams) Supply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamSupply])
}

type ImmutableMintParams struct {
	id int32
}

func (s ImmutableMintParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s ImmutableMintParams) To() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamTo])
}

type MutableMintParams struct {
	id int32
}

func (s MutableMintParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s MutableMintParams) To() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamTo])
}

type ImmutableTransferParams struct {
	id int32
}

func (s ImmutableTransferParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAccount])
}

func (s ImmutableTransferParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamAmount])
}

type MutableTransferParams struct {
	id int32
}

func (s MutableTransferParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAccount])
}

func (s MutableTransferParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamAmount])
}

type ImmutableTransferFromParams struct {
	id int32
}

func (s ImmutableTransferFromParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAccount])
}

func (s ImmutableTransferFromParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s ImmutableTransferFromParams) Recipient() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamRecipient])
}

type MutableTransferFromParams struct {
	id int32
}

func (s MutableTransferFromParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAccount])
}

func (s MutableTransferFromParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s MutableTransferFromParams) Recipient() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamRecipient])
}

type ImmutableAllowanceParams struct {
	id int32
}

func (s ImmutableAllowanceParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAccount])
}

func (s ImmutableAllowanceParams) Delegation() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamDelegation])
}

type MutableAllowanceParams struct {
	id int32
}

func (s MutableAllowanceParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAccount])
}

func (s MutableAllowanceParams) Delegation() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamDelegation])
}

type ImmutableBalanceOfParams struct {
	id int32
}

func (s ImmutableBalanceOfParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAccount])
}

type MutableBalanceOfParams struct {
	id int32
}

func (s MutableBalanceOfParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAccount])
}
