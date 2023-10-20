package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushPaperformatResult 结构体
type AlibabaDamaiMevOpenPushPaperformatResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushPaperformatResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushPaperformatResult)
	},
}

// GetAlibabaDamaiMevOpenPushPaperformatResult() 从对象池中获取AlibabaDamaiMevOpenPushPaperformatResult
func GetAlibabaDamaiMevOpenPushPaperformatResult() *AlibabaDamaiMevOpenPushPaperformatResult {
	return poolAlibabaDamaiMevOpenPushPaperformatResult.Get().(*AlibabaDamaiMevOpenPushPaperformatResult)
}

// ReleaseAlibabaDamaiMevOpenPushPaperformatResult 释放AlibabaDamaiMevOpenPushPaperformatResult
func ReleaseAlibabaDamaiMevOpenPushPaperformatResult(v *AlibabaDamaiMevOpenPushPaperformatResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushPaperformatResult.Put(v)
}
