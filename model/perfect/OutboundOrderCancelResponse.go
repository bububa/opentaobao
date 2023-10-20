package perfect

import (
	"sync"
)

// OutboundOrderCancelResponse 结构体
type OutboundOrderCancelResponse struct {
	// 1
	OutboundOrderCancels []OutboundOrderCancelDto `json:"outbound_order_cancels,omitempty" xml:"outbound_order_cancels>outbound_order_cancel_dto,omitempty"`
	// 1
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOutboundOrderCancelResponse = sync.Pool{
	New: func() any {
		return new(OutboundOrderCancelResponse)
	},
}

// GetOutboundOrderCancelResponse() 从对象池中获取OutboundOrderCancelResponse
func GetOutboundOrderCancelResponse() *OutboundOrderCancelResponse {
	return poolOutboundOrderCancelResponse.Get().(*OutboundOrderCancelResponse)
}

// ReleaseOutboundOrderCancelResponse 释放OutboundOrderCancelResponse
func ReleaseOutboundOrderCancelResponse(v *OutboundOrderCancelResponse) {
	v.OutboundOrderCancels = v.OutboundOrderCancels[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolOutboundOrderCancelResponse.Put(v)
}
