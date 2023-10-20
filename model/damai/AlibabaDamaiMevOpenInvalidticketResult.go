package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenInvalidticketResult 结构体
type AlibabaDamaiMevOpenInvalidticketResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenInvalidticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenInvalidticketResult)
	},
}

// GetAlibabaDamaiMevOpenInvalidticketResult() 从对象池中获取AlibabaDamaiMevOpenInvalidticketResult
func GetAlibabaDamaiMevOpenInvalidticketResult() *AlibabaDamaiMevOpenInvalidticketResult {
	return poolAlibabaDamaiMevOpenInvalidticketResult.Get().(*AlibabaDamaiMevOpenInvalidticketResult)
}

// ReleaseAlibabaDamaiMevOpenInvalidticketResult 释放AlibabaDamaiMevOpenInvalidticketResult
func ReleaseAlibabaDamaiMevOpenInvalidticketResult(v *AlibabaDamaiMevOpenInvalidticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenInvalidticketResult.Put(v)
}
