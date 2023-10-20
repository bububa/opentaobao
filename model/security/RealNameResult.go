package security

import (
	"sync"
)

// RealNameResult 结构体
type RealNameResult struct {
	// checkCode
	CheckCode string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	// checkMessage
	CheckMessage string `json:"check_message,omitempty" xml:"check_message,omitempty"`
	// match
	Match bool `json:"match,omitempty" xml:"match,omitempty"`
}

var poolRealNameResult = sync.Pool{
	New: func() any {
		return new(RealNameResult)
	},
}

// GetRealNameResult() 从对象池中获取RealNameResult
func GetRealNameResult() *RealNameResult {
	return poolRealNameResult.Get().(*RealNameResult)
}

// ReleaseRealNameResult 释放RealNameResult
func ReleaseRealNameResult(v *RealNameResult) {
	v.CheckCode = ""
	v.CheckMessage = ""
	v.Match = false
	poolRealNameResult.Put(v)
}
