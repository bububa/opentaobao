package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushvenueResult 结构体
type AlibabaDamaiMevOpenPushvenueResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 出参
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushvenueResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushvenueResult)
	},
}

// GetAlibabaDamaiMevOpenPushvenueResult() 从对象池中获取AlibabaDamaiMevOpenPushvenueResult
func GetAlibabaDamaiMevOpenPushvenueResult() *AlibabaDamaiMevOpenPushvenueResult {
	return poolAlibabaDamaiMevOpenPushvenueResult.Get().(*AlibabaDamaiMevOpenPushvenueResult)
}

// ReleaseAlibabaDamaiMevOpenPushvenueResult 释放AlibabaDamaiMevOpenPushvenueResult
func ReleaseAlibabaDamaiMevOpenPushvenueResult(v *AlibabaDamaiMevOpenPushvenueResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushvenueResult.Put(v)
}
