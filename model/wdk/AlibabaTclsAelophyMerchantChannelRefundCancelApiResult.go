package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelRefundCancelApiResult 结构体
type AlibabaTclsAelophyMerchantChannelRefundCancelApiResult struct {
	// 返回码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelRefundCancelApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelRefundCancelApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundCancelApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelRefundCancelApiResult
func GetAlibabaTclsAelophyMerchantChannelRefundCancelApiResult() *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult {
	return poolAlibabaTclsAelophyMerchantChannelRefundCancelApiResult.Get().(*AlibabaTclsAelophyMerchantChannelRefundCancelApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundCancelApiResult 释放AlibabaTclsAelophyMerchantChannelRefundCancelApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelRefundCancelApiResult(v *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelRefundCancelApiResult.Put(v)
}
