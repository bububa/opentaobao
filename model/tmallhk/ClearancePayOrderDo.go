package tmallhk

import (
	"sync"
)

// ClearancePayOrderDo 结构体
type ClearancePayOrderDo struct {
	// 支付宝买家ID
	AlipayBuyerId string `json:"alipay_buyer_id,omitempty" xml:"alipay_buyer_id,omitempty"`
	// 支付单号
	PayOrderId int64 `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
}

var poolClearancePayOrderDo = sync.Pool{
	New: func() any {
		return new(ClearancePayOrderDo)
	},
}

// GetClearancePayOrderDo() 从对象池中获取ClearancePayOrderDo
func GetClearancePayOrderDo() *ClearancePayOrderDo {
	return poolClearancePayOrderDo.Get().(*ClearancePayOrderDo)
}

// ReleaseClearancePayOrderDo 释放ClearancePayOrderDo
func ReleaseClearancePayOrderDo(v *ClearancePayOrderDo) {
	v.AlipayBuyerId = ""
	v.PayOrderId = 0
	poolClearancePayOrderDo.Put(v)
}
