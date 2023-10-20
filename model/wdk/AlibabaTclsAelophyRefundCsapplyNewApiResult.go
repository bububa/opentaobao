package wdk

import (
	"sync"
)

// AlibabaTclsAelophyRefundCsapplyNewApiResult 结构体
type AlibabaTclsAelophyRefundCsapplyNewApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyRefundCsapplyNewApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundCsapplyNewApiResult)
	},
}

// GetAlibabaTclsAelophyRefundCsapplyNewApiResult() 从对象池中获取AlibabaTclsAelophyRefundCsapplyNewApiResult
func GetAlibabaTclsAelophyRefundCsapplyNewApiResult() *AlibabaTclsAelophyRefundCsapplyNewApiResult {
	return poolAlibabaTclsAelophyRefundCsapplyNewApiResult.Get().(*AlibabaTclsAelophyRefundCsapplyNewApiResult)
}

// ReleaseAlibabaTclsAelophyRefundCsapplyNewApiResult 释放AlibabaTclsAelophyRefundCsapplyNewApiResult
func ReleaseAlibabaTclsAelophyRefundCsapplyNewApiResult(v *AlibabaTclsAelophyRefundCsapplyNewApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyRefundCsapplyNewApiResult.Put(v)
}
