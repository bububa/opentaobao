package dt

import (
	"sync"
)

// TribeError 结构体
type TribeError struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误说明
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
	v.Code = ""
	v.Message = ""
	poolTribeError.Put(v)
}
