package alsc

import (
	"sync"
)

// TribeError 结构体
type TribeError struct {
	// 错误原因
	View string `json:"view,omitempty" xml:"view,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTribeError = sync.Pool{
	New: func() any {
		return new(TribeError)
	},
}

// GetTribeError() 从对象池中获取TribeError
func GetTribeError() *TribeError {
	return poolTribeError.Get().(*TribeError)
}

// ReleaseTribeError 释放TribeError
func ReleaseTribeError(v *TribeError) {
	v.View = ""
	v.Code = ""
	v.Message = ""
	poolTribeError.Put(v)
}
