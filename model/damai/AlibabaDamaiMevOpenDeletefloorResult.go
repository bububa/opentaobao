package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeletefloorResult 结构体
type AlibabaDamaiMevOpenDeletefloorResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeletefloorResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletefloorResult)
	},
}

// GetAlibabaDamaiMevOpenDeletefloorResult() 从对象池中获取AlibabaDamaiMevOpenDeletefloorResult
func GetAlibabaDamaiMevOpenDeletefloorResult() *AlibabaDamaiMevOpenDeletefloorResult {
	return poolAlibabaDamaiMevOpenDeletefloorResult.Get().(*AlibabaDamaiMevOpenDeletefloorResult)
}

// ReleaseAlibabaDamaiMevOpenDeletefloorResult 释放AlibabaDamaiMevOpenDeletefloorResult
func ReleaseAlibabaDamaiMevOpenDeletefloorResult(v *AlibabaDamaiMevOpenDeletefloorResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeletefloorResult.Put(v)
}
