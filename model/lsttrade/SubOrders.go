package lsttrade

import (
	"sync"
)

// SubOrders 结构体
type SubOrders struct {
	// 退款数量
	RefundCount int64 `json:"refund_count,omitempty" xml:"refund_count,omitempty"`
	// 子单ID
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolSubOrders = sync.Pool{
	New: func() any {
		return new(SubOrders)
	},
}

// GetSubOrders() 从对象池中获取SubOrders
func GetSubOrders() *SubOrders {
	return poolSubOrders.Get().(*SubOrders)
}

// ReleaseSubOrders 释放SubOrders
func ReleaseSubOrders(v *SubOrders) {
	v.RefundCount = 0
	v.SubOrderId = 0
	poolSubOrders.Put(v)
}
