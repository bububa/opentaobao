package user

import (
	"sync"
)

// OpenAccountTokenApplyResult 结构体
type OpenAccountTokenApplyResult struct {
	// token
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

var poolOpenAccountTokenApplyResult = sync.Pool{
	New: func() any {
		return new(OpenAccountTokenApplyResult)
	},
}

// GetOpenAccountTokenApplyResult() 从对象池中获取OpenAccountTokenApplyResult
func GetOpenAccountTokenApplyResult() *OpenAccountTokenApplyResult {
	return poolOpenAccountTokenApplyResult.Get().(*OpenAccountTokenApplyResult)
}

// ReleaseOpenAccountTokenApplyResult 释放OpenAccountTokenApplyResult
func ReleaseOpenAccountTokenApplyResult(v *OpenAccountTokenApplyResult) {
	v.Data = ""
	v.Message = ""
	v.Code = 0
	v.Successful = false
	poolOpenAccountTokenApplyResult.Put(v)
}
