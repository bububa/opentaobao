package wdk

import (
	"sync"
)

// AlibabaTclsAelophyRefundFetchgoodsApiResult 结构体
type AlibabaTclsAelophyRefundFetchgoodsApiResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaTclsAelophyRefundFetchgoodsApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundFetchgoodsApiResult)
	},
}

// GetAlibabaTclsAelophyRefundFetchgoodsApiResult() 从对象池中获取AlibabaTclsAelophyRefundFetchgoodsApiResult
func GetAlibabaTclsAelophyRefundFetchgoodsApiResult() *AlibabaTclsAelophyRefundFetchgoodsApiResult {
	return poolAlibabaTclsAelophyRefundFetchgoodsApiResult.Get().(*AlibabaTclsAelophyRefundFetchgoodsApiResult)
}

// ReleaseAlibabaTclsAelophyRefundFetchgoodsApiResult 释放AlibabaTclsAelophyRefundFetchgoodsApiResult
func ReleaseAlibabaTclsAelophyRefundFetchgoodsApiResult(v *AlibabaTclsAelophyRefundFetchgoodsApiResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.IsSuccess = false
	poolAlibabaTclsAelophyRefundFetchgoodsApiResult.Put(v)
}
