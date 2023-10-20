package logistic

import (
	"sync"
)

// TopServiceResult 结构体
type TopServiceResult struct {
	// 接口返回数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 详细描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopServiceResult = sync.Pool{
	New: func() any {
		return new(TopServiceResult)
	},
}

// GetTopServiceResult() 从对象池中获取TopServiceResult
func GetTopServiceResult() *TopServiceResult {
	return poolTopServiceResult.Get().(*TopServiceResult)
}

// ReleaseTopServiceResult 释放TopServiceResult
func ReleaseTopServiceResult(v *TopServiceResult) {
	v.Data = ""
	v.Code = ""
	v.Msg = ""
	v.Description = ""
	v.Success = false
	poolTopServiceResult.Put(v)
}
