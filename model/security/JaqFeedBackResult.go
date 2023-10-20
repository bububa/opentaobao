package security

import (
	"sync"
)

// JaqFeedBackResult 结构体
type JaqFeedBackResult struct {
	// feedBack返回信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolJaqFeedBackResult = sync.Pool{
	New: func() any {
		return new(JaqFeedBackResult)
	},
}

// GetJaqFeedBackResult() 从对象池中获取JaqFeedBackResult
func GetJaqFeedBackResult() *JaqFeedBackResult {
	return poolJaqFeedBackResult.Get().(*JaqFeedBackResult)
}

// ReleaseJaqFeedBackResult 释放JaqFeedBackResult
func ReleaseJaqFeedBackResult(v *JaqFeedBackResult) {
	v.ErrorMsg = ""
	poolJaqFeedBackResult.Put(v)
}
