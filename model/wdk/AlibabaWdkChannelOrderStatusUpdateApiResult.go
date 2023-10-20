package wdk

import (
	"sync"
)

// AlibabaWdkChannelOrderStatusUpdateApiResult 结构体
type AlibabaWdkChannelOrderStatusUpdateApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回内容
	Model *OrderOperateResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkChannelOrderStatusUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderStatusUpdateApiResult)
	},
}

// GetAlibabaWdkChannelOrderStatusUpdateApiResult() 从对象池中获取AlibabaWdkChannelOrderStatusUpdateApiResult
func GetAlibabaWdkChannelOrderStatusUpdateApiResult() *AlibabaWdkChannelOrderStatusUpdateApiResult {
	return poolAlibabaWdkChannelOrderStatusUpdateApiResult.Get().(*AlibabaWdkChannelOrderStatusUpdateApiResult)
}

// ReleaseAlibabaWdkChannelOrderStatusUpdateApiResult 释放AlibabaWdkChannelOrderStatusUpdateApiResult
func ReleaseAlibabaWdkChannelOrderStatusUpdateApiResult(v *AlibabaWdkChannelOrderStatusUpdateApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkChannelOrderStatusUpdateApiResult.Put(v)
}
