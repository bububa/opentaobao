package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantUserUploadApiResult 结构体
type AlibabaTclsAelophyMerchantUserUploadApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaTclsAelophyMerchantUserUploadApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantUserUploadApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantUserUploadApiResult() 从对象池中获取AlibabaTclsAelophyMerchantUserUploadApiResult
func GetAlibabaTclsAelophyMerchantUserUploadApiResult() *AlibabaTclsAelophyMerchantUserUploadApiResult {
	return poolAlibabaTclsAelophyMerchantUserUploadApiResult.Get().(*AlibabaTclsAelophyMerchantUserUploadApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantUserUploadApiResult 释放AlibabaTclsAelophyMerchantUserUploadApiResult
func ReleaseAlibabaTclsAelophyMerchantUserUploadApiResult(v *AlibabaTclsAelophyMerchantUserUploadApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	v.Model = false
	poolAlibabaTclsAelophyMerchantUserUploadApiResult.Put(v)
}
