package sungari

import (
	"sync"
)

// TaobaoCloudbridgeCaseinvestExecuteResult 结构体
type TaobaoCloudbridgeCaseinvestExecuteResult struct {
	// data值，JSON数据，可转换成对应的结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 说明
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 接口调用是否成功，1：成功；2：失败
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoCloudbridgeCaseinvestExecuteResult = sync.Pool{
	New: func() any {
		return new(TaobaoCloudbridgeCaseinvestExecuteResult)
	},
}

// GetTaobaoCloudbridgeCaseinvestExecuteResult() 从对象池中获取TaobaoCloudbridgeCaseinvestExecuteResult
func GetTaobaoCloudbridgeCaseinvestExecuteResult() *TaobaoCloudbridgeCaseinvestExecuteResult {
	return poolTaobaoCloudbridgeCaseinvestExecuteResult.Get().(*TaobaoCloudbridgeCaseinvestExecuteResult)
}

// ReleaseTaobaoCloudbridgeCaseinvestExecuteResult 释放TaobaoCloudbridgeCaseinvestExecuteResult
func ReleaseTaobaoCloudbridgeCaseinvestExecuteResult(v *TaobaoCloudbridgeCaseinvestExecuteResult) {
	v.Data = ""
	v.Message = ""
	v.Code = 0
	poolTaobaoCloudbridgeCaseinvestExecuteResult.Put(v)
}
