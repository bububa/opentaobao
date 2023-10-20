package ascp

import (
	"sync"
)

// SubSourceOrder 结构体
type SubSourceOrder struct {
	// 交易平台子订单编码
	SubSourceOrderCode string `json:"sub_source_order_code,omitempty" xml:"sub_source_order_code,omitempty"`
}

var poolSubSourceOrder = sync.Pool{
	New: func() any {
		return new(SubSourceOrder)
	},
}

// GetSubSourceOrder() 从对象池中获取SubSourceOrder
func GetSubSourceOrder() *SubSourceOrder {
	return poolSubSourceOrder.Get().(*SubSourceOrder)
}

// ReleaseSubSourceOrder 释放SubSourceOrder
func ReleaseSubSourceOrder(v *SubSourceOrder) {
	v.SubSourceOrderCode = ""
	poolSubSourceOrder.Put(v)
}
