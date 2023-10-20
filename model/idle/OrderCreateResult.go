package idle

import (
	"sync"
)

// OrderCreateResult 结构体
type OrderCreateResult struct {
	// 订单编号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
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
	v.BizOrderId = 0
	poolOrderCreateResult.Put(v)
}
