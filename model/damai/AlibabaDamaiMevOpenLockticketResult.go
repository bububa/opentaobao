package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenLockticketResult 结构体
type AlibabaDamaiMevOpenLockticketResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenLockticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenLockticketResult)
	},
}

// GetAlibabaDamaiMevOpenLockticketResult() 从对象池中获取AlibabaDamaiMevOpenLockticketResult
func GetAlibabaDamaiMevOpenLockticketResult() *AlibabaDamaiMevOpenLockticketResult {
	return poolAlibabaDamaiMevOpenLockticketResult.Get().(*AlibabaDamaiMevOpenLockticketResult)
}

// ReleaseAlibabaDamaiMevOpenLockticketResult 释放AlibabaDamaiMevOpenLockticketResult
func ReleaseAlibabaDamaiMevOpenLockticketResult(v *AlibabaDamaiMevOpenLockticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenLockticketResult.Put(v)
}
