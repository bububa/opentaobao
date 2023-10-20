package paimai

import (
	"sync"
)

// Result4Top 结构体
type Result4Top struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 处理结果
	Value bool `json:"value,omitempty" xml:"value,omitempty"`
}

var poolResult4Top = sync.Pool{
	New: func() any {
		return new(Result4Top)
	},
}

// GetResult4Top() 从对象池中获取Result4Top
func GetResult4Top() *Result4Top {
	return poolResult4Top.Get().(*Result4Top)
}

// ReleaseResult4Top 释放Result4Top
func ReleaseResult4Top(v *Result4Top) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Value = false
	poolResult4Top.Put(v)
}
