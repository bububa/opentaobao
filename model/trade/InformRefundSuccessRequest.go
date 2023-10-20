package trade

import (
	"sync"
)

// InformRefundSuccessRequest 结构体
type InformRefundSuccessRequest struct {
	// 退货成功履约子单
	RefundFulfillSubOrders []RefundSuccessSubOrderRequest `json:"refund_fulfill_sub_orders,omitempty" xml:"refund_fulfill_sub_orders>refund_success_sub_order_request,omitempty"`
	// 门店编码
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolInformRefundSuccessRequest = sync.Pool{
	New: func() any {
		return new(InformRefundSuccessRequest)
	},
}

// GetInformRefundSuccessRequest() 从对象池中获取InformRefundSuccessRequest
func GetInformRefundSuccessRequest() *InformRefundSuccessRequest {
	return poolInformRefundSuccessRequest.Get().(*InformRefundSuccessRequest)
}

// ReleaseInformRefundSuccessRequest 释放InformRefundSuccessRequest
func ReleaseInformRefundSuccessRequest(v *InformRefundSuccessRequest) {
	v.RefundFulfillSubOrders = v.RefundFulfillSubOrders[:0]
	v.ShopId = ""
	poolInformRefundSuccessRequest.Put(v)
}
