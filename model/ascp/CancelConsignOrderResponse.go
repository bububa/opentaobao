package ascp

import (
	"sync"
)

// CancelConsignOrderResponse 结构体
type CancelConsignOrderResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCancelConsignOrderResponse = sync.Pool{
	New: func() any {
		return new(CancelConsignOrderResponse)
	},
}

// GetCancelConsignOrderResponse() 从对象池中获取CancelConsignOrderResponse
func GetCancelConsignOrderResponse() *CancelConsignOrderResponse {
	return poolCancelConsignOrderResponse.Get().(*CancelConsignOrderResponse)
}

// ReleaseCancelConsignOrderResponse 释放CancelConsignOrderResponse
func ReleaseCancelConsignOrderResponse(v *CancelConsignOrderResponse) {
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolCancelConsignOrderResponse.Put(v)
}
