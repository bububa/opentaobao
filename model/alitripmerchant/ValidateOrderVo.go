package alitripmerchant

import (
	"sync"
)

// ValidateOrderVo 结构体
type ValidateOrderVo struct {
	// 每日价格
	DailyPriceList []DailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>daily_price,omitempty"`
	// 价格变化描述
	AmountChangedDisplay string `json:"amount_changed_display,omitempty" xml:"amount_changed_display,omitempty"`
	// 总税价
	TotalTax string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
	// 实际需要支付金额，含税
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 价格发生变化
	IsAmountChanged bool `json:"is_amount_changed,omitempty" xml:"is_amount_changed,omitempty"`
}

var poolValidateOrderVo = sync.Pool{
	New: func() any {
		return new(ValidateOrderVo)
	},
}

// GetValidateOrderVo() 从对象池中获取ValidateOrderVo
func GetValidateOrderVo() *ValidateOrderVo {
	return poolValidateOrderVo.Get().(*ValidateOrderVo)
}

// ReleaseValidateOrderVo 释放ValidateOrderVo
func ReleaseValidateOrderVo(v *ValidateOrderVo) {
	v.DailyPriceList = v.DailyPriceList[:0]
	v.AmountChangedDisplay = ""
	v.TotalTax = ""
	v.TotalAmount = ""
	v.OrderCode = ""
	v.IsAmountChanged = false
	poolValidateOrderVo.Put(v)
}
