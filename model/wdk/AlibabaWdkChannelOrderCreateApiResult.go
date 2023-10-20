package wdk

import (
	"sync"
)

// AlibabaWdkChannelOrderCreateApiResult 结构体
type AlibabaWdkChannelOrderCreateApiResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回内容
	Model *OrderOperateResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkChannelOrderCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderCreateApiResult)
	},
}

// GetAlibabaWdkChannelOrderCreateApiResult() 从对象池中获取AlibabaWdkChannelOrderCreateApiResult
func GetAlibabaWdkChannelOrderCreateApiResult() *AlibabaWdkChannelOrderCreateApiResult {
	return poolAlibabaWdkChannelOrderCreateApiResult.Get().(*AlibabaWdkChannelOrderCreateApiResult)
}

// ReleaseAlibabaWdkChannelOrderCreateApiResult 释放AlibabaWdkChannelOrderCreateApiResult
func ReleaseAlibabaWdkChannelOrderCreateApiResult(v *AlibabaWdkChannelOrderCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkChannelOrderCreateApiResult.Put(v)
}
