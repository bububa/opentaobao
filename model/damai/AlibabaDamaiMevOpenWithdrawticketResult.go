package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenWithdrawticketResult 结构体
type AlibabaDamaiMevOpenWithdrawticketResult struct {
	// 错误内容
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenWithdrawticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenWithdrawticketResult)
	},
}

// GetAlibabaDamaiMevOpenWithdrawticketResult() 从对象池中获取AlibabaDamaiMevOpenWithdrawticketResult
func GetAlibabaDamaiMevOpenWithdrawticketResult() *AlibabaDamaiMevOpenWithdrawticketResult {
	return poolAlibabaDamaiMevOpenWithdrawticketResult.Get().(*AlibabaDamaiMevOpenWithdrawticketResult)
}

// ReleaseAlibabaDamaiMevOpenWithdrawticketResult 释放AlibabaDamaiMevOpenWithdrawticketResult
func ReleaseAlibabaDamaiMevOpenWithdrawticketResult(v *AlibabaDamaiMevOpenWithdrawticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenWithdrawticketResult.Put(v)
}
