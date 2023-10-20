package wdk

import (
	"sync"
)

// AlibabaWdkChannelCommentCreateApiResult 结构体
type AlibabaWdkChannelCommentCreateApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkChannelCommentCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelCommentCreateApiResult)
	},
}

// GetAlibabaWdkChannelCommentCreateApiResult() 从对象池中获取AlibabaWdkChannelCommentCreateApiResult
func GetAlibabaWdkChannelCommentCreateApiResult() *AlibabaWdkChannelCommentCreateApiResult {
	return poolAlibabaWdkChannelCommentCreateApiResult.Get().(*AlibabaWdkChannelCommentCreateApiResult)
}

// ReleaseAlibabaWdkChannelCommentCreateApiResult 释放AlibabaWdkChannelCommentCreateApiResult
func ReleaseAlibabaWdkChannelCommentCreateApiResult(v *AlibabaWdkChannelCommentCreateApiResult) {
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkChannelCommentCreateApiResult.Put(v)
}
