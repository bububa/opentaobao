package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenChangeticketResult 结构体
type AlibabaDamaiMevOpenChangeticketResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenChangeticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenChangeticketResult)
	},
}

// GetAlibabaDamaiMevOpenChangeticketResult() 从对象池中获取AlibabaDamaiMevOpenChangeticketResult
func GetAlibabaDamaiMevOpenChangeticketResult() *AlibabaDamaiMevOpenChangeticketResult {
	return poolAlibabaDamaiMevOpenChangeticketResult.Get().(*AlibabaDamaiMevOpenChangeticketResult)
}

// ReleaseAlibabaDamaiMevOpenChangeticketResult 释放AlibabaDamaiMevOpenChangeticketResult
func ReleaseAlibabaDamaiMevOpenChangeticketResult(v *AlibabaDamaiMevOpenChangeticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenChangeticketResult.Put(v)
}
