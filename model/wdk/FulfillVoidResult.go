package wdk

import (
	"sync"
)

// FulfillVoidResult 结构体
type FulfillVoidResult struct {
	// 返回码含义描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 返回码(success=false时不能为空)： SYSTEM_ERROR :系统异常(指令可重发) PARAM_ERROR :参数错误(指令不可重发，监控报警) BUSINESS_ERROR:业务异常(指令不可重发，监控报警)
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// true 调用成功 false 调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFulfillVoidResult = sync.Pool{
	New: func() any {
		return new(FulfillVoidResult)
	},
}

// GetFulfillVoidResult() 从对象池中获取FulfillVoidResult
func GetFulfillVoidResult() *FulfillVoidResult {
	return poolFulfillVoidResult.Get().(*FulfillVoidResult)
}

// ReleaseFulfillVoidResult 释放FulfillVoidResult
func ReleaseFulfillVoidResult(v *FulfillVoidResult) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolFulfillVoidResult.Put(v)
}
