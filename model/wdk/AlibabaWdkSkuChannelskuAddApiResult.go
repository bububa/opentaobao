package wdk

import (
	"sync"
)

// AlibabaWdkSkuChannelskuAddApiResult 结构体
type AlibabaWdkSkuChannelskuAddApiResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返会结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuChannelskuAddApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuAddApiResult)
	},
}

// GetAlibabaWdkSkuChannelskuAddApiResult() 从对象池中获取AlibabaWdkSkuChannelskuAddApiResult
func GetAlibabaWdkSkuChannelskuAddApiResult() *AlibabaWdkSkuChannelskuAddApiResult {
	return poolAlibabaWdkSkuChannelskuAddApiResult.Get().(*AlibabaWdkSkuChannelskuAddApiResult)
}

// ReleaseAlibabaWdkSkuChannelskuAddApiResult 释放AlibabaWdkSkuChannelskuAddApiResult
func ReleaseAlibabaWdkSkuChannelskuAddApiResult(v *AlibabaWdkSkuChannelskuAddApiResult) {
	v.ErrCode = ""
	v.Model = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuChannelskuAddApiResult.Put(v)
}
