package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushfaceResult 结构体
type AlibabaDamaiMevOpenPushfaceResult struct {
	// 错误内容
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushfaceResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushfaceResult)
	},
}

// GetAlibabaDamaiMevOpenPushfaceResult() 从对象池中获取AlibabaDamaiMevOpenPushfaceResult
func GetAlibabaDamaiMevOpenPushfaceResult() *AlibabaDamaiMevOpenPushfaceResult {
	return poolAlibabaDamaiMevOpenPushfaceResult.Get().(*AlibabaDamaiMevOpenPushfaceResult)
}

// ReleaseAlibabaDamaiMevOpenPushfaceResult 释放AlibabaDamaiMevOpenPushfaceResult
func ReleaseAlibabaDamaiMevOpenPushfaceResult(v *AlibabaDamaiMevOpenPushfaceResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushfaceResult.Put(v)
}
