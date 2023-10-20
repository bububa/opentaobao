package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushprojectResult 结构体
type AlibabaDamaiMevOpenPushprojectResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushprojectResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushprojectResult)
	},
}

// GetAlibabaDamaiMevOpenPushprojectResult() 从对象池中获取AlibabaDamaiMevOpenPushprojectResult
func GetAlibabaDamaiMevOpenPushprojectResult() *AlibabaDamaiMevOpenPushprojectResult {
	return poolAlibabaDamaiMevOpenPushprojectResult.Get().(*AlibabaDamaiMevOpenPushprojectResult)
}

// ReleaseAlibabaDamaiMevOpenPushprojectResult 释放AlibabaDamaiMevOpenPushprojectResult
func ReleaseAlibabaDamaiMevOpenPushprojectResult(v *AlibabaDamaiMevOpenPushprojectResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushprojectResult.Put(v)
}
