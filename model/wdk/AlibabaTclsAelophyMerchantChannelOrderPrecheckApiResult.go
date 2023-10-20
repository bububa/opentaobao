package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult 结构体
type AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult struct {
	// 返回码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 商品校验结果返回对象
	Model *PreCheckResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult
func GetAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult() *AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult {
	return poolAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult.Get().(*AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult 释放AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult(v *AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult.Put(v)
}
