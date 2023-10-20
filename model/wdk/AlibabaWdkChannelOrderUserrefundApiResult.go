package wdk

import (
	"sync"
)

// AlibabaWdkChannelOrderUserrefundApiResult 结构体
type AlibabaWdkChannelOrderUserrefundApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkChannelOrderUserrefundApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderUserrefundApiResult)
	},
}

// GetAlibabaWdkChannelOrderUserrefundApiResult() 从对象池中获取AlibabaWdkChannelOrderUserrefundApiResult
func GetAlibabaWdkChannelOrderUserrefundApiResult() *AlibabaWdkChannelOrderUserrefundApiResult {
	return poolAlibabaWdkChannelOrderUserrefundApiResult.Get().(*AlibabaWdkChannelOrderUserrefundApiResult)
}

// ReleaseAlibabaWdkChannelOrderUserrefundApiResult 释放AlibabaWdkChannelOrderUserrefundApiResult
func ReleaseAlibabaWdkChannelOrderUserrefundApiResult(v *AlibabaWdkChannelOrderUserrefundApiResult) {
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkChannelOrderUserrefundApiResult.Put(v)
}
