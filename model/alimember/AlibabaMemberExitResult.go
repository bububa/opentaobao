package alimember

import (
	"sync"
)

// AlibabaMemberExitResult 结构体
type AlibabaMemberExitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaMemberExitResult = sync.Pool{
	New: func() any {
		return new(AlibabaMemberExitResult)
	},
}

// GetAlibabaMemberExitResult() 从对象池中获取AlibabaMemberExitResult
func GetAlibabaMemberExitResult() *AlibabaMemberExitResult {
	return poolAlibabaMemberExitResult.Get().(*AlibabaMemberExitResult)
}

// ReleaseAlibabaMemberExitResult 释放AlibabaMemberExitResult
func ReleaseAlibabaMemberExitResult(v *AlibabaMemberExitResult) {
	v.Code = ""
	v.Message = ""
	poolAlibabaMemberExitResult.Put(v)
}
