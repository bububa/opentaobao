package wdk

import (
	"sync"
)

// ReverseResult 结构体
type ReverseResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口返回model
	Model *ApplyReverseResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolReverseResult = sync.Pool{
	New: func() any {
		return new(ReverseResult)
	},
}

// GetReverseResult() 从对象池中获取ReverseResult
func GetReverseResult() *ReverseResult {
	return poolReverseResult.Get().(*ReverseResult)
}

// ReleaseReverseResult 释放ReverseResult
func ReleaseReverseResult(v *ReverseResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Model = nil
	v.Success = false
	poolReverseResult.Put(v)
}
