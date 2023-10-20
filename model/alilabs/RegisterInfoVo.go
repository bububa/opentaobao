package alilabs

import (
	"sync"
)

// RegisterInfoVo 结构体
type RegisterInfoVo struct {
	// 用户开放id
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	// 设备uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolRegisterInfoVo = sync.Pool{
	New: func() any {
		return new(RegisterInfoVo)
	},
}

// GetRegisterInfoVo() 从对象池中获取RegisterInfoVo
func GetRegisterInfoVo() *RegisterInfoVo {
	return poolRegisterInfoVo.Get().(*RegisterInfoVo)
}

// ReleaseRegisterInfoVo 释放RegisterInfoVo
func ReleaseRegisterInfoVo(v *RegisterInfoVo) {
	v.UserOpenId = ""
	v.Uuid = ""
	poolRegisterInfoVo.Put(v)
}
