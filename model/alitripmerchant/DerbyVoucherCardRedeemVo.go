package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardRedeemVo 结构体
type DerbyVoucherCardRedeemVo struct {
	// 臻享卡卡号
	MemberVoucherCardID string `json:"member_voucher_card_i_d,omitempty" xml:"member_voucher_card_i_d,omitempty"`
	// 臻享卡卡类型
	VoucherCardCategory string `json:"voucher_card_category,omitempty" xml:"voucher_card_category,omitempty"`
	// 兑换成功
	RedeemSuccess bool `json:"redeem_success,omitempty" xml:"redeem_success,omitempty"`
}

var poolDerbyVoucherCardRedeemVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardRedeemVo)
	},
}

// GetDerbyVoucherCardRedeemVo() 从对象池中获取DerbyVoucherCardRedeemVo
func GetDerbyVoucherCardRedeemVo() *DerbyVoucherCardRedeemVo {
	return poolDerbyVoucherCardRedeemVo.Get().(*DerbyVoucherCardRedeemVo)
}

// ReleaseDerbyVoucherCardRedeemVo 释放DerbyVoucherCardRedeemVo
func ReleaseDerbyVoucherCardRedeemVo(v *DerbyVoucherCardRedeemVo) {
	v.MemberVoucherCardID = ""
	v.VoucherCardCategory = ""
	v.RedeemSuccess = false
	poolDerbyVoucherCardRedeemVo.Put(v)
}
