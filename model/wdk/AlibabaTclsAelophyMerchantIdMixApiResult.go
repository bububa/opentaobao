package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantIdMixApiResult 结构体
type AlibabaTclsAelophyMerchantIdMixApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回混淆id
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 获取mixId成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyMerchantIdMixApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantIdMixApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantIdMixApiResult() 从对象池中获取AlibabaTclsAelophyMerchantIdMixApiResult
func GetAlibabaTclsAelophyMerchantIdMixApiResult() *AlibabaTclsAelophyMerchantIdMixApiResult {
	return poolAlibabaTclsAelophyMerchantIdMixApiResult.Get().(*AlibabaTclsAelophyMerchantIdMixApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantIdMixApiResult 释放AlibabaTclsAelophyMerchantIdMixApiResult
func ReleaseAlibabaTclsAelophyMerchantIdMixApiResult(v *AlibabaTclsAelophyMerchantIdMixApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaTclsAelophyMerchantIdMixApiResult.Put(v)
}
