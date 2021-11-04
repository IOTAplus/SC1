// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use wasmlib::*;
use wasmlib::host::*;

pub struct PurchaseIEXP {
    pub amount:    i64,       // amount purchased
    pub error:     String,    // error to be reported to purchaser if anything goes wrong
    pub feedback:  String,    // the feedback for the person who purchase
    pub purchaser: ScAgentID, // who purchased
    pub timestamp: i64,       // when the purchase took place
}

impl PurchaseIEXP {
    pub fn from_bytes(bytes: &[u8]) -> PurchaseIEXP {
        let mut decode = BytesDecoder::new(bytes);
        PurchaseIEXP {
            amount: decode.int64(),
            error: decode.string(),
            feedback: decode.string(),
            purchaser: decode.agent_id(),
            timestamp: decode.int64(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.int64(self.amount);
        encode.string(&self.error);
        encode.string(&self.feedback);
        encode.agent_id(&self.purchaser);
        encode.int64(self.timestamp);
        return encode.data();
    }
}

pub struct ImmutablePurchaseIEXP {
    pub(crate) obj_id: i32,
    pub(crate) key_id: Key32,
}

impl ImmutablePurchaseIEXP {
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    pub fn value(&self) -> PurchaseIEXP {
        PurchaseIEXP::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_BYTES))
    }
}

pub struct MutablePurchaseIEXP {
    pub(crate) obj_id: i32,
    pub(crate) key_id: Key32,
}

impl MutablePurchaseIEXP {
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    pub fn set_value(&self, value: &PurchaseIEXP) {
        set_bytes(self.obj_id, self.key_id, TYPE_BYTES, &value.to_bytes());
    }

    pub fn value(&self) -> PurchaseIEXP {
        PurchaseIEXP::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_BYTES))
    }
}

// @formatter:on
