package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenBatchpushticketResult 结构体
type AlibabaDamaiMevOpenBatchpushticketResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenBatchpushticketResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenBatchpushticketResult)
	},
}

// GetAlibabaDamaiMevOpenBatchpushticketResult() 从对象池中获取AlibabaDamaiMevOpenBatchpushticketResult
func GetAlibabaDamaiMevOpenBatchpushticketResult() *AlibabaDamaiMevOpenBatchpushticketResult {
	return poolAlibabaDamaiMevOpenBatchpushticketResult.Get().(*AlibabaDamaiMevOpenBatchpushticketResult)
}

// ReleaseAlibabaDamaiMevOpenBatchpushticketResult 释放AlibabaDamaiMevOpenBatchpushticketResult
func ReleaseAlibabaDamaiMevOpenBatchpushticketResult(v *AlibabaDamaiMevOpenBatchpushticketResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenBatchpushticketResult.Put(v)
}
