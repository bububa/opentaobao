package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenUnlockticketResult 结构体
type AlibabaDamaiMevOpenUnlockticketResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenUnlockticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenUnlockticketResult)
	},
}

// GetAlibabaDamaiMevOpenUnlockticketResult() 从对象池中获取AlibabaDamaiMevOpenUnlockticketResult
func GetAlibabaDamaiMevOpenUnlockticketResult() *AlibabaDamaiMevOpenUnlockticketResult {
	return poolAlibabaDamaiMevOpenUnlockticketResult.Get().(*AlibabaDamaiMevOpenUnlockticketResult)
}

// ReleaseAlibabaDamaiMevOpenUnlockticketResult 释放AlibabaDamaiMevOpenUnlockticketResult
func ReleaseAlibabaDamaiMevOpenUnlockticketResult(v *AlibabaDamaiMevOpenUnlockticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenUnlockticketResult.Put(v)
}
