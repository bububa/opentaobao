package wdk

import (
	"sync"
)

// AlibabaTclsAelophyRefundDisagreeApiResult 结构体
type AlibabaTclsAelophyRefundDisagreeApiResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaTclsAelophyRefundDisagreeApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundDisagreeApiResult)
	},
}

// GetAlibabaTclsAelophyRefundDisagreeApiResult() 从对象池中获取AlibabaTclsAelophyRefundDisagreeApiResult
func GetAlibabaTclsAelophyRefundDisagreeApiResult() *AlibabaTclsAelophyRefundDisagreeApiResult {
	return poolAlibabaTclsAelophyRefundDisagreeApiResult.Get().(*AlibabaTclsAelophyRefundDisagreeApiResult)
}

// ReleaseAlibabaTclsAelophyRefundDisagreeApiResult 释放AlibabaTclsAelophyRefundDisagreeApiResult
func ReleaseAlibabaTclsAelophyRefundDisagreeApiResult(v *AlibabaTclsAelophyRefundDisagreeApiResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.IsSuccess = false
	poolAlibabaTclsAelophyRefundDisagreeApiResult.Put(v)
}
