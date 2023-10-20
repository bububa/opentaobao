package wdk

import (
	"sync"
)

// AlibabaWdkSkuChannelskuAddApiResults 结构体
type AlibabaWdkSkuChannelskuAddApiResults struct {
	// 返会结果集合
	Models []AlibabaWdkSkuChannelskuAddApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_channelsku_add_api_result,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuChannelskuAddApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuAddApiResults)
	},
}

// GetAlibabaWdkSkuChannelskuAddApiResults() 从对象池中获取AlibabaWdkSkuChannelskuAddApiResults
func GetAlibabaWdkSkuChannelskuAddApiResults() *AlibabaWdkSkuChannelskuAddApiResults {
	return poolAlibabaWdkSkuChannelskuAddApiResults.Get().(*AlibabaWdkSkuChannelskuAddApiResults)
}

// ReleaseAlibabaWdkSkuChannelskuAddApiResults 释放AlibabaWdkSkuChannelskuAddApiResults
func ReleaseAlibabaWdkSkuChannelskuAddApiResults(v *AlibabaWdkSkuChannelskuAddApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuChannelskuAddApiResults.Put(v)
}
