package mos

import (
	"sync"
)

// SingleResult 结构体
type SingleResult struct {
	// 系统错误
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 系统错误
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 成功返回
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolSingleResult = sync.Pool{
	New: func() any {
		return new(SingleResult)
	},
}

// GetSingleResult() 从对象池中获取SingleResult
func GetSingleResult() *SingleResult {
	return poolSingleResult.Get().(*SingleResult)
}

// ReleaseSingleResult 释放SingleResult
func ReleaseSingleResult(v *SingleResult) {
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Success = false
	v.Data = false
	poolSingleResult.Put(v)
}
