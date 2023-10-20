package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeleteitemResult 结构体
type AlibabaDamaiMevOpenDeleteitemResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeleteitemResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteitemResult)
	},
}

// GetAlibabaDamaiMevOpenDeleteitemResult() 从对象池中获取AlibabaDamaiMevOpenDeleteitemResult
func GetAlibabaDamaiMevOpenDeleteitemResult() *AlibabaDamaiMevOpenDeleteitemResult {
	return poolAlibabaDamaiMevOpenDeleteitemResult.Get().(*AlibabaDamaiMevOpenDeleteitemResult)
}

// ReleaseAlibabaDamaiMevOpenDeleteitemResult 释放AlibabaDamaiMevOpenDeleteitemResult
func ReleaseAlibabaDamaiMevOpenDeleteitemResult(v *AlibabaDamaiMevOpenDeleteitemResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeleteitemResult.Put(v)
}
