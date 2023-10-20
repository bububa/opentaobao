package tvpay

import (
	"sync"
)

// TopResultDo 结构体
type TopResultDo struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结构体
	Data *GetLoginInfoByOrderResultDo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopResultDo = sync.Pool{
	New: func() any {
		return new(TopResultDo)
	},
}

// GetTopResultDo() 从对象池中获取TopResultDo
func GetTopResultDo() *TopResultDo {
	return poolTopResultDo.Get().(*TopResultDo)
}

// ReleaseTopResultDo 释放TopResultDo
func ReleaseTopResultDo(v *TopResultDo) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolTopResultDo.Put(v)
}
