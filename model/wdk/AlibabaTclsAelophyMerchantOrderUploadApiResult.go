package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantOrderUploadApiResult 结构体
type AlibabaTclsAelophyMerchantOrderUploadApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaTclsAelophyMerchantOrderUploadApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantOrderUploadApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantOrderUploadApiResult() 从对象池中获取AlibabaTclsAelophyMerchantOrderUploadApiResult
func GetAlibabaTclsAelophyMerchantOrderUploadApiResult() *AlibabaTclsAelophyMerchantOrderUploadApiResult {
	return poolAlibabaTclsAelophyMerchantOrderUploadApiResult.Get().(*AlibabaTclsAelophyMerchantOrderUploadApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantOrderUploadApiResult 释放AlibabaTclsAelophyMerchantOrderUploadApiResult
func ReleaseAlibabaTclsAelophyMerchantOrderUploadApiResult(v *AlibabaTclsAelophyMerchantOrderUploadApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	v.Model = false
	poolAlibabaTclsAelophyMerchantOrderUploadApiResult.Put(v)
}
