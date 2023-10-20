package btrip

import (
	"sync"
)

// ClientInfoDo 结构体
type ClientInfoDo struct {
	// 乘机人id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 乘机人姓名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

var poolClientInfoDo = sync.Pool{
	New: func() any {
		return new(ClientInfoDo)
	},
}

// GetClientInfoDo() 从对象池中获取ClientInfoDo
func GetClientInfoDo() *ClientInfoDo {
	return poolClientInfoDo.Get().(*ClientInfoDo)
}

// ReleaseClientInfoDo 释放ClientInfoDo
func ReleaseClientInfoDo(v *ClientInfoDo) {
	v.UserId = ""
	v.UserName = ""
	poolClientInfoDo.Put(v)
}
