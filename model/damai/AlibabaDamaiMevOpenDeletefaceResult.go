package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeletefaceResult 结构体
type AlibabaDamaiMevOpenDeletefaceResult struct {
	// 错误内容
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeletefaceResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletefaceResult)
	},
}

// GetAlibabaDamaiMevOpenDeletefaceResult() 从对象池中获取AlibabaDamaiMevOpenDeletefaceResult
func GetAlibabaDamaiMevOpenDeletefaceResult() *AlibabaDamaiMevOpenDeletefaceResult {
	return poolAlibabaDamaiMevOpenDeletefaceResult.Get().(*AlibabaDamaiMevOpenDeletefaceResult)
}

// ReleaseAlibabaDamaiMevOpenDeletefaceResult 释放AlibabaDamaiMevOpenDeletefaceResult
func ReleaseAlibabaDamaiMevOpenDeletefaceResult(v *AlibabaDamaiMevOpenDeletefaceResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeletefaceResult.Put(v)
}
