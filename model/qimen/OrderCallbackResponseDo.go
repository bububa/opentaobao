package qimen

import (
	"sync"
)

// OrderCallbackResponseDo 结构体
type OrderCallbackResponseDo struct {
	// 响应结果:success|failure,success,string(10),必填,
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码,0,string(50),,
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息,invalid appkey,string(100),,
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolOrderCallbackResponseDo = sync.Pool{
	New: func() any {
		return new(OrderCallbackResponseDo)
	},
}

// GetOrderCallbackResponseDo() 从对象池中获取OrderCallbackResponseDo
func GetOrderCallbackResponseDo() *OrderCallbackResponseDo {
	return poolOrderCallbackResponseDo.Get().(*OrderCallbackResponseDo)
}

// ReleaseOrderCallbackResponseDo 释放OrderCallbackResponseDo
func ReleaseOrderCallbackResponseDo(v *OrderCallbackResponseDo) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolOrderCallbackResponseDo.Put(v)
}
