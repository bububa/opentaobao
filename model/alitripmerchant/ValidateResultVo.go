package alitripmerchant

import (
	"sync"
)

// ValidateResultVo 结构体
type ValidateResultVo struct {
	// 每日价格列表
	DailyPriceList []DailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>daily_price,omitempty"`
	// 每日价格列表
	DiningPolicyList []DailyInfo `json:"dining_policy_list,omitempty" xml:"dining_policy_list>daily_info,omitempty"`
	// 订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 价格变化
	AmountChangedDisplay string `json:"amount_changed_display,omitempty" xml:"amount_changed_display,omitempty"`
	// 总税价
	TotalTax string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
	// 总价不含税
	TotalPriceExcludeTax string `json:"total_price_exclude_tax,omitempty" xml:"total_price_exclude_tax,omitempty"`
	// 实际需要支付金额，含税
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 支付类型
	PaymentType string `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 取消政策
	CancelPolicy *CancelPolicy `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	// 加价信息
	MarkupInfo *MarkupInfoVo `json:"markup_info,omitempty" xml:"markup_info,omitempty"`
	// 外币信息
	ForeignCurrency *ForeignCurrencyDto `json:"foreign_currency,omitempty" xml:"foreign_currency,omitempty"`
	// 价格发生变化
	IsAmountChanged bool `json:"is_amount_changed,omitempty" xml:"is_amount_changed,omitempty"`
	// 是否为外币支付
	ForeignCurrencyPayment bool `json:"foreign_currency_payment,omitempty" xml:"foreign_currency_payment,omitempty"`
}

var poolValidateResultVo = sync.Pool{
	New: func() any {
		return new(ValidateResultVo)
	},
}

// GetValidateResultVo() 从对象池中获取ValidateResultVo
func GetValidateResultVo() *ValidateResultVo {
	return poolValidateResultVo.Get().(*ValidateResultVo)
}

// ReleaseValidateResultVo 释放ValidateResultVo
func ReleaseValidateResultVo(v *ValidateResultVo) {
	v.DailyPriceList = v.DailyPriceList[:0]
	v.DiningPolicyList = v.DiningPolicyList[:0]
	v.OrderCode = ""
	v.AmountChangedDisplay = ""
	v.TotalTax = ""
	v.TotalPriceExcludeTax = ""
	v.TotalAmount = ""
	v.PaymentType = ""
	v.CancelPolicy = nil
	v.MarkupInfo = nil
	v.ForeignCurrency = nil
	v.IsAmountChanged = false
	v.ForeignCurrencyPayment = false
	poolValidateResultVo.Put(v)
}
