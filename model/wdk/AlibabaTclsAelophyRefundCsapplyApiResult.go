package wdk

import (
	"sync"
)

// AlibabaTclsAelophyRefundCsapplyApiResult 结构体
type AlibabaTclsAelophyRefundCsapplyApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAelophyRefundCsapplyApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundCsapplyApiResult)
	},
}

// GetAlibabaTclsAelophyRefundCsapplyApiResult() 从对象池中获取AlibabaTclsAelophyRefundCsapplyApiResult
func GetAlibabaTclsAelophyRefundCsapplyApiResult() *AlibabaTclsAelophyRefundCsapplyApiResult {
	return poolAlibabaTclsAelophyRefundCsapplyApiResult.Get().(*AlibabaTclsAelophyRefundCsapplyApiResult)
}

// ReleaseAlibabaTclsAelophyRefundCsapplyApiResult 释放AlibabaTclsAelophyRefundCsapplyApiResult
func ReleaseAlibabaTclsAelophyRefundCsapplyApiResult(v *AlibabaTclsAelophyRefundCsapplyApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaTclsAelophyRefundCsapplyApiResult.Put(v)
}
