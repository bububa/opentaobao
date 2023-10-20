package aecreatives

import (
	"sync"
)

// ResponseResult 结构体
type ResponseResult struct {
	// 返回状态描述
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 返回状态码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 返回记录结果列表
	Result *AliexpressAffiliateCategoryGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

var poolResponseResult = sync.Pool{
	New: func() any {
		return new(ResponseResult)
	},
}

// GetResponseResult() 从对象池中获取ResponseResult
func GetResponseResult() *ResponseResult {
	return poolResponseResult.Get().(*ResponseResult)
}

// ReleaseResponseResult 释放ResponseResult
func ReleaseResponseResult(v *ResponseResult) {
	v.RespMsg = ""
	v.RespCode = 0
	v.Result = nil
	poolResponseResult.Put(v)
}
