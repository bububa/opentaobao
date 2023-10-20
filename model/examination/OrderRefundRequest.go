package examination

import (
	"sync"
)

// OrderRefundRequest 结构体
type OrderRefundRequest struct {
	// 子单信息
	SubOrderList []SubOrderRequest `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_request,omitempty"`
	// 主订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderRefundRequest = sync.Pool{
	New: func() any {
		return new(OrderRefundRequest)
	},
}

// GetOrderRefundRequest() 从对象池中获取OrderRefundRequest
func GetOrderRefundRequest() *OrderRefundRequest {
	return poolOrderRefundRequest.Get().(*OrderRefundRequest)
}

// ReleaseOrderRefundRequest 释放OrderRefundRequest
func ReleaseOrderRefundRequest(v *OrderRefundRequest) {
	v.SubOrderList = v.SubOrderList[:0]
	v.OrderId = ""
	poolOrderRefundRequest.Put(v)
}
