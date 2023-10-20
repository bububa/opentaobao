package wdk

import (
	"sync"
)

// AlibabaAxChannelSkuStatusUpdateApiResult 结构体
type AlibabaAxChannelSkuStatusUpdateApiResult struct {
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回错误bian ma
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口返回结果成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAxChannelSkuStatusUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaAxChannelSkuStatusUpdateApiResult)
	},
}

// GetAlibabaAxChannelSkuStatusUpdateApiResult() 从对象池中获取AlibabaAxChannelSkuStatusUpdateApiResult
func GetAlibabaAxChannelSkuStatusUpdateApiResult() *AlibabaAxChannelSkuStatusUpdateApiResult {
	return poolAlibabaAxChannelSkuStatusUpdateApiResult.Get().(*AlibabaAxChannelSkuStatusUpdateApiResult)
}

// ReleaseAlibabaAxChannelSkuStatusUpdateApiResult 释放AlibabaAxChannelSkuStatusUpdateApiResult
func ReleaseAlibabaAxChannelSkuStatusUpdateApiResult(v *AlibabaAxChannelSkuStatusUpdateApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAxChannelSkuStatusUpdateApiResult.Put(v)
}
