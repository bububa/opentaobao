package user

import (
	"sync"
)

// OpenaccountVoid 结构体
type OpenaccountVoid struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

var poolOpenaccountVoid = sync.Pool{
	New: func() any {
		return new(OpenaccountVoid)
	},
}

// GetOpenaccountVoid() 从对象池中获取OpenaccountVoid
func GetOpenaccountVoid() *OpenaccountVoid {
	return poolOpenaccountVoid.Get().(*OpenaccountVoid)
}

// ReleaseOpenaccountVoid 释放OpenaccountVoid
func ReleaseOpenaccountVoid(v *OpenaccountVoid) {
	v.Message = ""
	v.Code = 0
	v.Successful = false
	poolOpenaccountVoid.Put(v)
}
