package wdk

import (
	"sync"
)

// UmsResult 结构体
type UmsResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolUmsResult = sync.Pool{
	New: func() any {
		return new(UmsResult)
	},
}

// GetUmsResult() 从对象池中获取UmsResult
func GetUmsResult() *UmsResult {
	return poolUmsResult.Get().(*UmsResult)
}

// ReleaseUmsResult 释放UmsResult
func ReleaseUmsResult(v *UmsResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolUmsResult.Put(v)
}
