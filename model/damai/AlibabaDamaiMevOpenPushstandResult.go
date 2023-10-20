package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushstandResult 结构体
type AlibabaDamaiMevOpenPushstandResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushstandResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushstandResult)
	},
}

// GetAlibabaDamaiMevOpenPushstandResult() 从对象池中获取AlibabaDamaiMevOpenPushstandResult
func GetAlibabaDamaiMevOpenPushstandResult() *AlibabaDamaiMevOpenPushstandResult {
	return poolAlibabaDamaiMevOpenPushstandResult.Get().(*AlibabaDamaiMevOpenPushstandResult)
}

// ReleaseAlibabaDamaiMevOpenPushstandResult 释放AlibabaDamaiMevOpenPushstandResult
func ReleaseAlibabaDamaiMevOpenPushstandResult(v *AlibabaDamaiMevOpenPushstandResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushstandResult.Put(v)
}
