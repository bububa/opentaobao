package idle

import (
	"sync"
)

// RecycleRefundTopRequest 结构体
type RecycleRefundTopRequest struct {
	// 申请仅退款
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolRecycleRefundTopRequest = sync.Pool{
	New: func() any {
		return new(RecycleRefundTopRequest)
	},
}

// GetRecycleRefundTopRequest() 从对象池中获取RecycleRefundTopRequest
func GetRecycleRefundTopRequest() *RecycleRefundTopRequest {
	return poolRecycleRefundTopRequest.Get().(*RecycleRefundTopRequest)
}

// ReleaseRecycleRefundTopRequest 释放RecycleRefundTopRequest
func ReleaseRecycleRefundTopRequest(v *RecycleRefundTopRequest) {
	v.Message = ""
	v.BizOrderId = 0
	poolRecycleRefundTopRequest.Put(v)
}
