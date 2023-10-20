package wdk

import (
	"sync"
)

// AlibabaTclsAelophyMerchantOrderBatchUploadApiResult 结构体
type AlibabaTclsAelophyMerchantOrderBatchUploadApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaTclsAelophyMerchantOrderBatchUploadApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantOrderBatchUploadApiResult)
	},
}

// GetAlibabaTclsAelophyMerchantOrderBatchUploadApiResult() 从对象池中获取AlibabaTclsAelophyMerchantOrderBatchUploadApiResult
func GetAlibabaTclsAelophyMerchantOrderBatchUploadApiResult() *AlibabaTclsAelophyMerchantOrderBatchUploadApiResult {
	return poolAlibabaTclsAelophyMerchantOrderBatchUploadApiResult.Get().(*AlibabaTclsAelophyMerchantOrderBatchUploadApiResult)
}

// ReleaseAlibabaTclsAelophyMerchantOrderBatchUploadApiResult 释放AlibabaTclsAelophyMerchantOrderBatchUploadApiResult
func ReleaseAlibabaTclsAelophyMerchantOrderBatchUploadApiResult(v *AlibabaTclsAelophyMerchantOrderBatchUploadApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	v.Model = false
	poolAlibabaTclsAelophyMerchantOrderBatchUploadApiResult.Put(v)
}
