// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_init(ctx: &ScFuncContext, f: &InitContext) {
    if f.params.owner().exists() {
        f.state.owner().set_value(&f.params.owner().value());
        return;
    }
    f.state.owner().set_value(&ctx.contract_creator());
}

pub fn func_set_owner(_ctx: &ScFuncContext, _f: &SetOwnerContext) {
}

pub fn view_get_owner(_ctx: &ScViewContext, _f: &GetOwnerContext) {
}

pub fn func_purchase(_ctx: &ScFuncContext, _f: &PurchaseContext) {
}

pub fn func_withdraw(_ctx: &ScFuncContext, _f: &WithdrawContext) {
}

pub fn view_purchase_info(_ctx: &ScViewContext, _f: &PurchaseInfoContext) {
}

pub fn view_purchase_view(_ctx: &ScViewContext, _f: &PurchaseViewContext) {
}
