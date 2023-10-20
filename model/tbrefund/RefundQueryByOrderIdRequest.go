package tbrefund

import (
	"sync"
)

// RefundQueryByOrderIdRequest 结构体
type RefundQueryByOrderIdRequest struct {
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolRefundQueryByOrderIdRequest = sync.Pool{
	New: func() any {
		return new(RefundQueryByOrderIdRequest)
	},
}

// GetRefundQueryByOrderIdRequest() 从对象池中获取RefundQueryByOrderIdRequest
func GetRefundQueryByOrderIdRequest() *RefundQueryByOrderIdRequest {
	return poolRefundQueryByOrderIdRequest.Get().(*RefundQueryByOrderIdRequest)
}

// ReleaseRefundQueryByOrderIdRequest 释放RefundQueryByOrderIdRequest
func ReleaseRefundQueryByOrderIdRequest(v *RefundQueryByOrderIdRequest) {
	v.BizOrderId = 0
	poolRefundQueryByOrderIdRequest.Put(v)
}
