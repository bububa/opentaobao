package damaiticklet

import (
	"sync"
)

// OpenResult 结构体
type OpenResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolOpenResult = sync.Pool{
	New: func() any {
		return new(OpenResult)
	},
}

// GetOpenResult() 从对象池中获取OpenResult
func GetOpenResult() *OpenResult {
	return poolOpenResult.Get().(*OpenResult)
}

// ReleaseOpenResult 释放OpenResult
func ReleaseOpenResult(v *OpenResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Model = false
	poolOpenResult.Put(v)
}
