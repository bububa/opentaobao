package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushfaceelementResult 结构体
type AlibabaDamaiMevOpenPushfaceelementResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushfaceelementResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushfaceelementResult)
	},
}

// GetAlibabaDamaiMevOpenPushfaceelementResult() 从对象池中获取AlibabaDamaiMevOpenPushfaceelementResult
func GetAlibabaDamaiMevOpenPushfaceelementResult() *AlibabaDamaiMevOpenPushfaceelementResult {
	return poolAlibabaDamaiMevOpenPushfaceelementResult.Get().(*AlibabaDamaiMevOpenPushfaceelementResult)
}

// ReleaseAlibabaDamaiMevOpenPushfaceelementResult 释放AlibabaDamaiMevOpenPushfaceelementResult
func ReleaseAlibabaDamaiMevOpenPushfaceelementResult(v *AlibabaDamaiMevOpenPushfaceelementResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushfaceelementResult.Put(v)
}
