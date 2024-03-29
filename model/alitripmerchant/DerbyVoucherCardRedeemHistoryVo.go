package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardRedeemHistoryVo 结构体
type DerbyVoucherCardRedeemHistoryVo struct {
	// 臻享卡卡号
	MemberVoucherCardID string `json:"member_voucher_card_i_d,omitempty" xml:"member_voucher_card_i_d,omitempty"`
	// 会员卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 臻享卡兑换码
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
	// 臻享卡类型
	VoucherCardCategory string `json:"voucher_card_category,omitempty" xml:"voucher_card_category,omitempty"`
	// 兑换时间
	RedeemTime string `json:"redeem_time,omitempty" xml:"redeem_time,omitempty"`
	// 兑换成功
	RedeemSuccess bool `json:"redeem_success,omitempty" xml:"redeem_success,omitempty"`
}

var poolDerbyVoucherCardRedeemHistoryVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardRedeemHistoryVo)
	},
}

// GetDerbyVoucherCardRedeemHistoryVo() 从对象池中获取DerbyVoucherCardRedeemHistoryVo
func GetDerbyVoucherCardRedeemHistoryVo() *DerbyVoucherCardRedeemHistoryVo {
	return poolDerbyVoucherCardRedeemHistoryVo.Get().(*DerbyVoucherCardRedeemHistoryVo)
}

// ReleaseDerbyVoucherCardRedeemHistoryVo 释放DerbyVoucherCardRedeemHistoryVo
func ReleaseDerbyVoucherCardRedeemHistoryVo(v *DerbyVoucherCardRedeemHistoryVo) {
	v.MemberVoucherCardID = ""
	v.CardNumber = ""
	v.TicketCode = ""
	v.VoucherCardCategory = ""
	v.RedeemTime = ""
	v.RedeemSuccess = false
	poolDerbyVoucherCardRedeemHistoryVo.Put(v)
}
