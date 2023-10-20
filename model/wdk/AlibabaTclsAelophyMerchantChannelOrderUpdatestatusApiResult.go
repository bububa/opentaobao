package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult 结构体
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult
func GetAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult() *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult {
	return poolAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult.Get().(*AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult 释放AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult(v *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult.Put(v)
}
