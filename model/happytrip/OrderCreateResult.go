package happytrip

import (
	"sync"
)

// OrderCreateResult 结构体
type OrderCreateResult struct {
	// 订单信息
	Order *OrderInfo `json:"order,omitempty" xml:"order,omitempty"`
	// 价格信息
	Price *PriceInfo `json:"price,omitempty" xml:"price,omitempty"`
}

var poolOrderCreateResult = sync.Pool{
	New: func() any {
		return new(OrderCreateResult)
	},
}

// GetOrderCreateResult() 从对象池中获取OrderCreateResult
func GetOrderCreateResult() *OrderCreateResult {
	return poolOrderCreateResult.Get().(*OrderCreateResult)
}

// ReleaseOrderCreateResult 释放OrderCreateResult
func ReleaseOrderCreateResult(v *OrderCreateResult) {
	v.Order = nil
	v.Price = nil
	poolOrderCreateResult.Put(v)
}
