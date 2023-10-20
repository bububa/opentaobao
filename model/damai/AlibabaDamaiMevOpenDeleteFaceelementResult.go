package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenDeleteFaceelementResult 结构体
type AlibabaDamaiMevOpenDeleteFaceelementResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenDeleteFaceelementResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteFaceelementResult)
	},
}

// GetAlibabaDamaiMevOpenDeleteFaceelementResult() 从对象池中获取AlibabaDamaiMevOpenDeleteFaceelementResult
func GetAlibabaDamaiMevOpenDeleteFaceelementResult() *AlibabaDamaiMevOpenDeleteFaceelementResult {
	return poolAlibabaDamaiMevOpenDeleteFaceelementResult.Get().(*AlibabaDamaiMevOpenDeleteFaceelementResult)
}

// ReleaseAlibabaDamaiMevOpenDeleteFaceelementResult 释放AlibabaDamaiMevOpenDeleteFaceelementResult
func ReleaseAlibabaDamaiMevOpenDeleteFaceelementResult(v *AlibabaDamaiMevOpenDeleteFaceelementResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenDeleteFaceelementResult.Put(v)
}
