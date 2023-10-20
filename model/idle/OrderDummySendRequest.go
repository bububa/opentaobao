package idle

import (
	"sync"
)

// OrderDummySendRequest 结构体
type OrderDummySendRequest struct {
	// 需要虚拟发货的订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolOrderDummySendRequest = sync.Pool{
	New: func() any {
		return new(OrderDummySendRequest)
	},
}

// GetOrderDummySendRequest() 从对象池中获取OrderDummySendRequest
func GetOrderDummySendRequest() *OrderDummySendRequest {
	return poolOrderDummySendRequest.Get().(*OrderDummySendRequest)
}

// ReleaseOrderDummySendRequest 释放OrderDummySendRequest
func ReleaseOrderDummySendRequest(v *OrderDummySendRequest) {
	v.BizOrderId = 0
	poolOrderDummySendRequest.Put(v)
}
