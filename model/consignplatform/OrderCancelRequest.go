package consignplatform

import (
	"sync"
)

// OrderCancelRequest 结构体
type OrderCancelRequest struct {
	// 外部订单id列表
	SubOuterOrderIds []string `json:"sub_outer_order_ids,omitempty" xml:"sub_outer_order_ids>string,omitempty"`
	// 外部订单id
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 订单来源
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
}

var poolOrderCancelRequest = sync.Pool{
	New: func() any {
		return new(OrderCancelRequest)
	},
}

// GetOrderCancelRequest() 从对象池中获取OrderCancelRequest
func GetOrderCancelRequest() *OrderCancelRequest {
	return poolOrderCancelRequest.Get().(*OrderCancelRequest)
}

// ReleaseOrderCancelRequest 释放OrderCancelRequest
func ReleaseOrderCancelRequest(v *OrderCancelRequest) {
	v.SubOuterOrderIds = v.SubOuterOrderIds[:0]
	v.OuterOrderId = ""
	v.OrderSource = ""
	poolOrderCancelRequest.Put(v)
}
