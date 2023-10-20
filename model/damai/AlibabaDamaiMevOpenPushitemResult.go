package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushitemResult 结构体
type AlibabaDamaiMevOpenPushitemResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushitemResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushitemResult)
	},
}

// GetAlibabaDamaiMevOpenPushitemResult() 从对象池中获取AlibabaDamaiMevOpenPushitemResult
func GetAlibabaDamaiMevOpenPushitemResult() *AlibabaDamaiMevOpenPushitemResult {
	return poolAlibabaDamaiMevOpenPushitemResult.Get().(*AlibabaDamaiMevOpenPushitemResult)
}

// ReleaseAlibabaDamaiMevOpenPushitemResult 释放AlibabaDamaiMevOpenPushitemResult
func ReleaseAlibabaDamaiMevOpenPushitemResult(v *AlibabaDamaiMevOpenPushitemResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushitemResult.Put(v)
}
