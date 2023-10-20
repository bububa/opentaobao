package tmallhk

import (
	"sync"
)

// OrderCertRequest 结构体
type OrderCertRequest struct {
	// 订单编号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderCertRequest = sync.Pool{
	New: func() any {
		return new(OrderCertRequest)
	},
}

// GetOrderCertRequest() 从对象池中获取OrderCertRequest
func GetOrderCertRequest() *OrderCertRequest {
	return poolOrderCertRequest.Get().(*OrderCertRequest)
}

// ReleaseOrderCertRequest 释放OrderCertRequest
func ReleaseOrderCertRequest(v *OrderCertRequest) {
	v.OrderId = 0
	poolOrderCertRequest.Put(v)
}
