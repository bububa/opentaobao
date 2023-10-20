package wdk

import (
	"sync"
)

// AlibabaWdkChannelOrderRefundConfirmApiResult 结构体
type AlibabaWdkChannelOrderRefundConfirmApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkChannelOrderRefundConfirmApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderRefundConfirmApiResult)
	},
}

// GetAlibabaWdkChannelOrderRefundConfirmApiResult() 从对象池中获取AlibabaWdkChannelOrderRefundConfirmApiResult
func GetAlibabaWdkChannelOrderRefundConfirmApiResult() *AlibabaWdkChannelOrderRefundConfirmApiResult {
	return poolAlibabaWdkChannelOrderRefundConfirmApiResult.Get().(*AlibabaWdkChannelOrderRefundConfirmApiResult)
}

// ReleaseAlibabaWdkChannelOrderRefundConfirmApiResult 释放AlibabaWdkChannelOrderRefundConfirmApiResult
func ReleaseAlibabaWdkChannelOrderRefundConfirmApiResult(v *AlibabaWdkChannelOrderRefundConfirmApiResult) {
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkChannelOrderRefundConfirmApiResult.Put(v)
}
