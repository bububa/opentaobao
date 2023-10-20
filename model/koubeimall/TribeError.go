package koubeimall

import (
	"sync"
)

// TribeError 结构体
type TribeError struct {
	// 错误信息可读性描述
	View string `json:"view,omitempty" xml:"view,omitempty"`
	// 错误信息描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
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
	v.Message = ""
	v.Code = ""
	poolTribeError.Put(v)
}
