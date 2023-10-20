package examination

import (
	"sync"
)

// SubOrderRequest 结构体
type SubOrderRequest struct {
	// 子单id
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 子单退款备注
	RefundNote string `json:"refund_note,omitempty" xml:"refund_note,omitempty"`
	// 子单退款金额
	RefundAmount string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

var poolSubOrderRequest = sync.Pool{
	New: func() any {
		return new(SubOrderRequest)
	},
}

// GetSubOrderRequest() 从对象池中获取SubOrderRequest
func GetSubOrderRequest() *SubOrderRequest {
	return poolSubOrderRequest.Get().(*SubOrderRequest)
}

// ReleaseSubOrderRequest 释放SubOrderRequest
func ReleaseSubOrderRequest(v *SubOrderRequest) {
	v.SubOrderId = ""
	v.RefundNote = ""
	v.RefundAmount = ""
	poolSubOrderRequest.Put(v)
}
