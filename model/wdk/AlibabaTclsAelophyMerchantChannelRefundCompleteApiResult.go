package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult 结构体
type AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult struct {
	// 返回码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult
func GetAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult() *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult {
	return poolAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult.Get().(*AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult 释放AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult(v *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelRefundCompleteApiResult.Put(v)
}
