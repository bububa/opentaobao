package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelRefundApplyApiResult 结构体
type AlibabaTclsAelophyMerchantChannelRefundApplyApiResult struct {
	// 返回码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelRefundApplyApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelRefundApplyApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundApplyApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelRefundApplyApiResult
func GetAlibabaTclsAelophyMerchantChannelRefundApplyApiResult() *AlibabaTclsAelophyMerchantChannelRefundApplyApiResult {
	return poolAlibabaTclsAelophyMerchantChannelRefundApplyApiResult.Get().(*AlibabaTclsAelophyMerchantChannelRefundApplyApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundApplyApiResult 释放AlibabaTclsAelophyMerchantChannelRefundApplyApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelRefundApplyApiResult(v *AlibabaTclsAelophyMerchantChannelRefundApplyApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelRefundApplyApiResult.Put(v)
}
