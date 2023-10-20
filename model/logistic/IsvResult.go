package logistic

import (
	"sync"
)

// IsvResult 结构体
type IsvResult struct {
	// 共享码
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ServerErrorCode string `json:"server_error_code,omitempty" xml:"server_error_code,omitempty"`
	// 描述
	Describe string `json:"describe,omitempty" xml:"describe,omitempty"`
}

var poolIsvResult = sync.Pool{
	New: func() any {
		return new(IsvResult)
	},
}

// GetIsvResult() 从对象池中获取IsvResult
func GetIsvResult() *IsvResult {
	return poolIsvResult.Get().(*IsvResult)
}

// ReleaseIsvResult 释放IsvResult
func ReleaseIsvResult(v *IsvResult) {
	v.Data = ""
	v.ServerErrorCode = ""
	v.Describe = ""
	poolIsvResult.Put(v)
}
