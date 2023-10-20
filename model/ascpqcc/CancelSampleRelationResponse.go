package ascpqcc

import (
	"sync"
)

// CancelSampleRelationResponse 结构体
type CancelSampleRelationResponse struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功标示
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCancelSampleRelationResponse = sync.Pool{
	New: func() any {
		return new(CancelSampleRelationResponse)
	},
}

// GetCancelSampleRelationResponse() 从对象池中获取CancelSampleRelationResponse
func GetCancelSampleRelationResponse() *CancelSampleRelationResponse {
	return poolCancelSampleRelationResponse.Get().(*CancelSampleRelationResponse)
}

// ReleaseCancelSampleRelationResponse 释放CancelSampleRelationResponse
func ReleaseCancelSampleRelationResponse(v *CancelSampleRelationResponse) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolCancelSampleRelationResponse.Put(v)
}
