package wdk

import (
	"sync"
)

// AlibabaTclsAelophyRefundAgreeApiResult 结构体
type AlibabaTclsAelophyRefundAgreeApiResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaTclsAelophyRefundAgreeApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundAgreeApiResult)
	},
}

// GetAlibabaTclsAelophyRefundAgreeApiResult() 从对象池中获取AlibabaTclsAelophyRefundAgreeApiResult
func GetAlibabaTclsAelophyRefundAgreeApiResult() *AlibabaTclsAelophyRefundAgreeApiResult {
	return poolAlibabaTclsAelophyRefundAgreeApiResult.Get().(*AlibabaTclsAelophyRefundAgreeApiResult)
}

// ReleaseAlibabaTclsAelophyRefundAgreeApiResult 释放AlibabaTclsAelophyRefundAgreeApiResult
func ReleaseAlibabaTclsAelophyRefundAgreeApiResult(v *AlibabaTclsAelophyRefundAgreeApiResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.IsSuccess = false
	poolAlibabaTclsAelophyRefundAgreeApiResult.Put(v)
}
