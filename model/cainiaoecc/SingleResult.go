package cainiaoecc

import (
	"sync"
)

// SingleResult 结构体
type SingleResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 数据对象
	Result *DelayExceptionCountDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSingleResult = sync.Pool{
	New: func() any {
		return new(SingleResult)
	},
}

// GetSingleResult() 从对象池中获取SingleResult
func GetSingleResult() *SingleResult {
	return poolSingleResult.Get().(*SingleResult)
}

// ReleaseSingleResult 释放SingleResult
func ReleaseSingleResult(v *SingleResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Result = nil
	v.IsRetry = false
	v.Success = false
	poolSingleResult.Put(v)
}
