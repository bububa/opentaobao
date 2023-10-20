package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeleteprojectResult 结构体
type AlibabaDamaiMevOpenDeleteprojectResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeleteprojectResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteprojectResult)
	},
}

// GetAlibabaDamaiMevOpenDeleteprojectResult() 从对象池中获取AlibabaDamaiMevOpenDeleteprojectResult
func GetAlibabaDamaiMevOpenDeleteprojectResult() *AlibabaDamaiMevOpenDeleteprojectResult {
	return poolAlibabaDamaiMevOpenDeleteprojectResult.Get().(*AlibabaDamaiMevOpenDeleteprojectResult)
}

// ReleaseAlibabaDamaiMevOpenDeleteprojectResult 释放AlibabaDamaiMevOpenDeleteprojectResult
func ReleaseAlibabaDamaiMevOpenDeleteprojectResult(v *AlibabaDamaiMevOpenDeleteprojectResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeleteprojectResult.Put(v)
}
