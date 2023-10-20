package wdk

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolBaseResult.Put(v)
}
