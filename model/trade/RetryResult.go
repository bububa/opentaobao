package trade

import (
	"sync"
)

// RetryResult 结构体
type RetryResult struct {
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolRetryResult = sync.Pool{
	New: func() any {
		return new(RetryResult)
	},
}

// GetRetryResult() 从对象池中获取RetryResult
func GetRetryResult() *RetryResult {
	return poolRetryResult.Get().(*RetryResult)
}

// ReleaseRetryResult 释放RetryResult
func ReleaseRetryResult(v *RetryResult) {
	v.ExtInfo = ""
	v.Error = nil
	v.Success = false
	v.Retry = false
	poolRetryResult.Put(v)
}
