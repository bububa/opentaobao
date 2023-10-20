package drug

import (
	"sync"
)

// ResponseResult 结构体
type ResponseResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果
	Result *PageResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	v.Success = false
	poolResponseResult.Put(v)
}
