package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantChannelOrderCreateApiResult 结构体
type AlibabaTclsAelophyMerchantChannelOrderCreateApiResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantChannelOrderCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderCreateApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderCreateApiResult() 从对象池中获取AlibabaTclsAelophyMerchantChannelOrderCreateApiResult
func GetAlibabaTclsAelophyMerchantChannelOrderCreateApiResult() *AlibabaTclsAelophyMerchantChannelOrderCreateApiResult {
	return poolAlibabaTclsAelophyMerchantChannelOrderCreateApiResult.Get().(*AlibabaTclsAelophyMerchantChannelOrderCreateApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderCreateApiResult 释放AlibabaTclsAelophyMerchantChannelOrderCreateApiResult
func ReleaseAlibabaTclsAelophyMerchantChannelOrderCreateApiResult(v *AlibabaTclsAelophyMerchantChannelOrderCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantChannelOrderCreateApiResult.Put(v)
}
