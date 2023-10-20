package iot

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 响应code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 返回结果
	RetValue *BusinessRecipeOpenDto `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.Message = ""
	v.TraceId = ""
	v.RetCode = 0
	v.RetValue = nil
	poolBaseResult.Put(v)
}
