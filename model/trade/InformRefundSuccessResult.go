package trade

import (
	"sync"
)

// InformRefundSuccessResult 结构体
type InformRefundSuccessResult struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolInformRefundSuccessResult = sync.Pool{
	New: func() any {
		return new(InformRefundSuccessResult)
	},
}

// GetInformRefundSuccessResult() 从对象池中获取InformRefundSuccessResult
func GetInformRefundSuccessResult() *InformRefundSuccessResult {
	return poolInformRefundSuccessResult.Get().(*InformRefundSuccessResult)
}

// ReleaseInformRefundSuccessResult 释放InformRefundSuccessResult
func ReleaseInformRefundSuccessResult(v *InformRefundSuccessResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolInformRefundSuccessResult.Put(v)
}
