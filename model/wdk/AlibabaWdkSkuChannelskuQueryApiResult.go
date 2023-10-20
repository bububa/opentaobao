package wdk

import (
	"sync"
)

// AlibabaWdkSkuChannelskuQueryApiResult 结构体
type AlibabaWdkSkuChannelskuQueryApiResult struct {
	// 异常状态码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 业务数据模型
	Model *ChannelSkuDo `json:"model,omitempty" xml:"model,omitempty"`
	// 业务调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuChannelskuQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuQueryApiResult)
	},
}

// GetAlibabaWdkSkuChannelskuQueryApiResult() 从对象池中获取AlibabaWdkSkuChannelskuQueryApiResult
func GetAlibabaWdkSkuChannelskuQueryApiResult() *AlibabaWdkSkuChannelskuQueryApiResult {
	return poolAlibabaWdkSkuChannelskuQueryApiResult.Get().(*AlibabaWdkSkuChannelskuQueryApiResult)
}

// ReleaseAlibabaWdkSkuChannelskuQueryApiResult 释放AlibabaWdkSkuChannelskuQueryApiResult
func ReleaseAlibabaWdkSkuChannelskuQueryApiResult(v *AlibabaWdkSkuChannelskuQueryApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkSkuChannelskuQueryApiResult.Put(v)
}
