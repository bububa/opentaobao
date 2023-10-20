package exchange

import (
	"sync"
)

// RefundBaseResponse 结构体
type RefundBaseResponse struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// exchange
	Exchange *Exchange `json:"exchange,omitempty" xml:"exchange,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRefundBaseResponse = sync.Pool{
	New: func() any {
		return new(RefundBaseResponse)
	},
}

// GetRefundBaseResponse() 从对象池中获取RefundBaseResponse
func GetRefundBaseResponse() *RefundBaseResponse {
	return poolRefundBaseResponse.Get().(*RefundBaseResponse)
}

// ReleaseRefundBaseResponse 释放RefundBaseResponse
func ReleaseRefundBaseResponse(v *RefundBaseResponse) {
	v.Message = ""
	v.MsgCode = ""
	v.Exchange = nil
	v.Success = false
	poolRefundBaseResponse.Put(v)
}
