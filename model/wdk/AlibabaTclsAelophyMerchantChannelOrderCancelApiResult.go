package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelOrderCancelApiResult 结构体
type AlibabaTclsAelophyMerchantChannelOrderCancelApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelOrderCancelApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderCancelApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderCancelApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelOrderCancelApiResult
func GetAlibabaTclsAelophyMerchantChannelOrderCancelApiResult() *AlibabaTclsAelophyMerchantChannelOrderCancelApiResult {
	return poolAlibabaTclsAelophyMerchantChannelOrderCancelApiResult.Get().(*AlibabaTclsAelophyMerchantChannelOrderCancelApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderCancelApiResult 释放AlibabaTclsAelophyMerchantChannelOrderCancelApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelOrderCancelApiResult(v *AlibabaTclsAelophyMerchantChannelOrderCancelApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelOrderCancelApiResult.Put(v)
}
