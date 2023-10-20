package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardActiveVo 结构体
type DerbyVoucherCardActiveVo struct {
	// 会员卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 权益卡号
	MemberVoucherCardID string `json:"member_voucher_card_i_d,omitempty" xml:"member_voucher_card_i_d,omitempty"`
}

var poolDerbyVoucherCardActiveVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardActiveVo)
	},
}

// GetDerbyVoucherCardActiveVo() 从对象池中获取DerbyVoucherCardActiveVo
func GetDerbyVoucherCardActiveVo() *DerbyVoucherCardActiveVo {
	return poolDerbyVoucherCardActiveVo.Get().(*DerbyVoucherCardActiveVo)
}

// ReleaseDerbyVoucherCardActiveVo 释放DerbyVoucherCardActiveVo
func ReleaseDerbyVoucherCardActiveVo(v *DerbyVoucherCardActiveVo) {
	v.CardNumber = ""
	v.MemberVoucherCardID = ""
	poolDerbyVoucherCardActiveVo.Put(v)
}
