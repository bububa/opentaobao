package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushfloorResult 结构体
type AlibabaDamaiMevOpenPushfloorResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushfloorResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushfloorResult)
	},
}

// GetAlibabaDamaiMevOpenPushfloorResult() 从对象池中获取AlibabaDamaiMevOpenPushfloorResult
func GetAlibabaDamaiMevOpenPushfloorResult() *AlibabaDamaiMevOpenPushfloorResult {
	return poolAlibabaDamaiMevOpenPushfloorResult.Get().(*AlibabaDamaiMevOpenPushfloorResult)
}

// ReleaseAlibabaDamaiMevOpenPushfloorResult 释放AlibabaDamaiMevOpenPushfloorResult
func ReleaseAlibabaDamaiMevOpenPushfloorResult(v *AlibabaDamaiMevOpenPushfloorResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushfloorResult.Put(v)
}
