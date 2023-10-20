package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeleteperformResult 结构体
type AlibabaDamaiMevOpenDeleteperformResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeleteperformResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteperformResult)
	},
}

// GetAlibabaDamaiMevOpenDeleteperformResult() 从对象池中获取AlibabaDamaiMevOpenDeleteperformResult
func GetAlibabaDamaiMevOpenDeleteperformResult() *AlibabaDamaiMevOpenDeleteperformResult {
	return poolAlibabaDamaiMevOpenDeleteperformResult.Get().(*AlibabaDamaiMevOpenDeleteperformResult)
}

// ReleaseAlibabaDamaiMevOpenDeleteperformResult 释放AlibabaDamaiMevOpenDeleteperformResult
func ReleaseAlibabaDamaiMevOpenDeleteperformResult(v *AlibabaDamaiMevOpenDeleteperformResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeleteperformResult.Put(v)
}
