package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeletestandResult 结构体
type AlibabaDamaiMevOpenDeletestandResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeletestandResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletestandResult)
	},
}

// GetAlibabaDamaiMevOpenDeletestandResult() 从对象池中获取AlibabaDamaiMevOpenDeletestandResult
func GetAlibabaDamaiMevOpenDeletestandResult() *AlibabaDamaiMevOpenDeletestandResult {
	return poolAlibabaDamaiMevOpenDeletestandResult.Get().(*AlibabaDamaiMevOpenDeletestandResult)
}

// ReleaseAlibabaDamaiMevOpenDeletestandResult 释放AlibabaDamaiMevOpenDeletestandResult
func ReleaseAlibabaDamaiMevOpenDeletestandResult(v *AlibabaDamaiMevOpenDeletestandResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeletestandResult.Put(v)
}
