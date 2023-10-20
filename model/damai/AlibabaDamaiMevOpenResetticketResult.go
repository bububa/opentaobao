package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenResetticketResult 结构体
type AlibabaDamaiMevOpenResetticketResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenResetticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenResetticketResult)
	},
}

// GetAlibabaDamaiMevOpenResetticketResult() 从对象池中获取AlibabaDamaiMevOpenResetticketResult
func GetAlibabaDamaiMevOpenResetticketResult() *AlibabaDamaiMevOpenResetticketResult {
	return poolAlibabaDamaiMevOpenResetticketResult.Get().(*AlibabaDamaiMevOpenResetticketResult)
}

// ReleaseAlibabaDamaiMevOpenResetticketResult 释放AlibabaDamaiMevOpenResetticketResult
func ReleaseAlibabaDamaiMevOpenResetticketResult(v *AlibabaDamaiMevOpenResetticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenResetticketResult.Put(v)
}
