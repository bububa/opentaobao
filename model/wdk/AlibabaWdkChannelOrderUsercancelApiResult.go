package wdk

import (
	"sync"
)

// AlibabaWdkChannelOrderUsercancelApiResult 结构体
type AlibabaWdkChannelOrderUsercancelApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkChannelOrderUsercancelApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderUsercancelApiResult)
	},
}

// GetAlibabaWdkChannelOrderUsercancelApiResult() 从对象池中获取AlibabaWdkChannelOrderUsercancelApiResult
func GetAlibabaWdkChannelOrderUsercancelApiResult() *AlibabaWdkChannelOrderUsercancelApiResult {
	return poolAlibabaWdkChannelOrderUsercancelApiResult.Get().(*AlibabaWdkChannelOrderUsercancelApiResult)
}

// ReleaseAlibabaWdkChannelOrderUsercancelApiResult 释放AlibabaWdkChannelOrderUsercancelApiResult
func ReleaseAlibabaWdkChannelOrderUsercancelApiResult(v *AlibabaWdkChannelOrderUsercancelApiResult) {
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkChannelOrderUsercancelApiResult.Put(v)
}
