package wdk

import (
	"sync"
)

// AlibabaWdkSkuChannelskuUpdateApiResult 结构体
type AlibabaWdkSkuChannelskuUpdateApiResult struct {
	// 单个商品
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuChannelskuUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuUpdateApiResult)
	},
}

// GetAlibabaWdkSkuChannelskuUpdateApiResult() 从对象池中获取AlibabaWdkSkuChannelskuUpdateApiResult
func GetAlibabaWdkSkuChannelskuUpdateApiResult() *AlibabaWdkSkuChannelskuUpdateApiResult {
	return poolAlibabaWdkSkuChannelskuUpdateApiResult.Get().(*AlibabaWdkSkuChannelskuUpdateApiResult)
}

// ReleaseAlibabaWdkSkuChannelskuUpdateApiResult 释放AlibabaWdkSkuChannelskuUpdateApiResult
func ReleaseAlibabaWdkSkuChannelskuUpdateApiResult(v *AlibabaWdkSkuChannelskuUpdateApiResult) {
	v.Model = ""
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaWdkSkuChannelskuUpdateApiResult.Put(v)
}
