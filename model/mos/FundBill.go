package mos

import (
	"sync"
)

// FundBill 结构体
type FundBill struct {
	// 业务扩展参数，json格式
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 子资金渠道。必然返回。取值为支付宝红包：alipay_coupon 支付宝余额：alipay_account 支付宝集分宝：alipay_point 支付宝余额宝：alipay_financeaccount 蚂蚁花呗：alipay_pcredit 支付宝预付卡：alipay_pcard 喵街储值卡（经支付宝打款）：alipay_mj_vcard 喵街购物券（经支付宝打款）：alipay_mj_voucher 喵街补贴（经支付宝打款）：alipay_mj_subsidy 其它（经支付宝打款）：alipay_other
	SubFundChannel string `json:"sub_fund_channel,omitempty" xml:"sub_fund_channel,omitempty"`
	// 打款类型。必然返回。取值为alipay，代表打款类型为支付宝（打款到商户的支付宝账号）
	FundChannel string `json:"fund_channel,omitempty" xml:"fund_channel,omitempty"`
	// 金额。单位为人民币（分）。必然返回
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolFundBill = sync.Pool{
	New: func() any {
		return new(FundBill)
	},
}

// GetFundBill() 从对象池中获取FundBill
func GetFundBill() *FundBill {
	return poolFundBill.Get().(*FundBill)
}

// ReleaseFundBill 释放FundBill
func ReleaseFundBill(v *FundBill) {
	v.ExtendParams = ""
	v.SubFundChannel = ""
	v.FundChannel = ""
	v.Amount = 0
	poolFundBill.Put(v)
}
