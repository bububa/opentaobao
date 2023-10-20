package ascpqcc

import (
	"sync"
)

// SendResult 结构体
type SendResult struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 业务系统返回响应
	CancelResponse *CancelSampleRelationResponse `json:"cancel_response,omitempty" xml:"cancel_response,omitempty"`
	// 业务系统返回结果
	UpdateResponse *UpdateSampleResponse `json:"update_response,omitempty" xml:"update_response,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSendResult = sync.Pool{
	New: func() any {
		return new(SendResult)
	},
}

// GetSendResult() 从对象池中获取SendResult
func GetSendResult() *SendResult {
	return poolSendResult.Get().(*SendResult)
}

// ReleaseSendResult 释放SendResult
func ReleaseSendResult(v *SendResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.CancelResponse = nil
	v.UpdateResponse = nil
	v.Success = false
	poolSendResult.Put(v)
}
