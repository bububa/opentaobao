package happytrip

import (
	"sync"
)

// OrderConfirmResult 结构体
type OrderConfirmResult struct {
	// 返回的id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderConfirmResult = sync.Pool{
	New: func() any {
		return new(OrderConfirmResult)
	},
}

// GetOrderConfirmResult() 从对象池中获取OrderConfirmResult
func GetOrderConfirmResult() *OrderConfirmResult {
	return poolOrderConfirmResult.Get().(*OrderConfirmResult)
}

// ReleaseOrderConfirmResult 释放OrderConfirmResult
func ReleaseOrderConfirmResult(v *OrderConfirmResult) {
	v.OrderId = ""
	poolOrderConfirmResult.Put(v)
}
