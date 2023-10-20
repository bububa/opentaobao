package user

import (
	"sync"
)

// OpenaccountObject 结构体
type OpenaccountObject struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *OpenAccount `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

var poolOpenaccountObject = sync.Pool{
	New: func() any {
		return new(OpenaccountObject)
	},
}

// GetOpenaccountObject() 从对象池中获取OpenaccountObject
func GetOpenaccountObject() *OpenaccountObject {
	return poolOpenaccountObject.Get().(*OpenaccountObject)
}

// ReleaseOpenaccountObject 释放OpenaccountObject
func ReleaseOpenaccountObject(v *OpenaccountObject) {
	v.Message = ""
	v.Data = nil
	v.Code = 0
	v.Successful = false
	poolOpenaccountObject.Put(v)
}
