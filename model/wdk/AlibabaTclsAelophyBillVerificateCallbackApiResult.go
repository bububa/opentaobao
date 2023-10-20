package wdk

import (
	"sync"
)

// AlibabaTclsAelophyBillVerificateCallbackApiResult 结构体
type AlibabaTclsAelophyBillVerificateCallbackApiResult struct {
	// 错误说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 回调是否处理成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyBillVerificateCallbackApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyBillVerificateCallbackApiResult)
	},
}

// GetAlibabaTclsAelophyBillVerificateCallbackApiResult() 从对象池中获取AlibabaTclsAelophyBillVerificateCallbackApiResult
func GetAlibabaTclsAelophyBillVerificateCallbackApiResult() *AlibabaTclsAelophyBillVerificateCallbackApiResult {
	return poolAlibabaTclsAelophyBillVerificateCallbackApiResult.Get().(*AlibabaTclsAelophyBillVerificateCallbackApiResult)
}

// ReleaseAlibabaTclsAelophyBillVerificateCallbackApiResult 释放AlibabaTclsAelophyBillVerificateCallbackApiResult
func ReleaseAlibabaTclsAelophyBillVerificateCallbackApiResult(v *AlibabaTclsAelophyBillVerificateCallbackApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyBillVerificateCallbackApiResult.Put(v)
}
