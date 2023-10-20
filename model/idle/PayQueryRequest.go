package idle

import (
	"sync"
)

// PayQueryRequest 结构体
type PayQueryRequest struct {
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolPayQueryRequest = sync.Pool{
	New: func() any {
		return new(PayQueryRequest)
	},
}

// GetPayQueryRequest() 从对象池中获取PayQueryRequest
func GetPayQueryRequest() *PayQueryRequest {
	return poolPayQueryRequest.Get().(*PayQueryRequest)
}

// ReleasePayQueryRequest 释放PayQueryRequest
func ReleasePayQueryRequest(v *PayQueryRequest) {
	v.BizOrderId = ""
	poolPayQueryRequest.Put(v)
}
