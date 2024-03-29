package wdk

import (
	"sync"
)

// AlibabaWdkSkuChannelskuQueryApiResults 结构体
type AlibabaWdkSkuChannelskuQueryApiResults struct {
	// 业务数据模型
	Models []AlibabaWdkSkuChannelskuQueryApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_channelsku_query_api_result,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// skuCode不能为空
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuChannelskuQueryApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuQueryApiResults)
	},
}

// GetAlibabaWdkSkuChannelskuQueryApiResults() 从对象池中获取AlibabaWdkSkuChannelskuQueryApiResults
func GetAlibabaWdkSkuChannelskuQueryApiResults() *AlibabaWdkSkuChannelskuQueryApiResults {
	return poolAlibabaWdkSkuChannelskuQueryApiResults.Get().(*AlibabaWdkSkuChannelskuQueryApiResults)
}

// ReleaseAlibabaWdkSkuChannelskuQueryApiResults 释放AlibabaWdkSkuChannelskuQueryApiResults
func ReleaseAlibabaWdkSkuChannelskuQueryApiResults(v *AlibabaWdkSkuChannelskuQueryApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuChannelskuQueryApiResults.Put(v)
}
