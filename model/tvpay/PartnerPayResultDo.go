package tvpay

import (
	"sync"
)

// PartnerPayResultDo 结构体
type PartnerPayResultDo struct {
	// 金额
	FundMoney string `json:"fund_money,omitempty" xml:"fund_money,omitempty"`
	// 金额构成
	FundMoneyCode string `json:"fund_money_code,omitempty" xml:"fund_money_code,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 支付结果码
	PayCode string `json:"pay_code,omitempty" xml:"pay_code,omitempty"`
	// 支付模式码
	PayMode string `json:"pay_mode,omitempty" xml:"pay_mode,omitempty"`
}

var poolPartnerPayResultDo = sync.Pool{
	New: func() any {
		return new(PartnerPayResultDo)
	},
}

// GetPartnerPayResultDo() 从对象池中获取PartnerPayResultDo
func GetPartnerPayResultDo() *PartnerPayResultDo {
	return poolPartnerPayResultDo.Get().(*PartnerPayResultDo)
}

// ReleasePartnerPayResultDo 释放PartnerPayResultDo
func ReleasePartnerPayResultDo(v *PartnerPayResultDo) {
	v.FundMoney = ""
	v.FundMoneyCode = ""
	v.Mobile = ""
	v.OrderNo = ""
	v.PayCode = ""
	v.PayMode = ""
	poolPartnerPayResultDo.Put(v)
}
