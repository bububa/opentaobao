package wdk

import (
	"sync"
)

// AlibabaWdkSkuChannelskuUpdateApiResults 结构体
type AlibabaWdkSkuChannelskuUpdateApiResults struct {
	// 单个商品返回结果集合
	Models []AlibabaWdkSkuChannelskuUpdateApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_channelsku_update_api_result,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuChannelskuUpdateApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuUpdateApiResults)
	},
}

// GetAlibabaWdkSkuChannelskuUpdateApiResults() 从对象池中获取AlibabaWdkSkuChannelskuUpdateApiResults
func GetAlibabaWdkSkuChannelskuUpdateApiResults() *AlibabaWdkSkuChannelskuUpdateApiResults {
	return poolAlibabaWdkSkuChannelskuUpdateApiResults.Get().(*AlibabaWdkSkuChannelskuUpdateApiResults)
}

// ReleaseAlibabaWdkSkuChannelskuUpdateApiResults 释放AlibabaWdkSkuChannelskuUpdateApiResults
func ReleaseAlibabaWdkSkuChannelskuUpdateApiResults(v *AlibabaWdkSkuChannelskuUpdateApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuChannelskuUpdateApiResults.Put(v)
}
