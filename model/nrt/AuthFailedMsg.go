package nrt

import (
	"sync"
)

// AuthFailedMsg 结构体
type AuthFailedMsg struct {
	// 入参手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 添加失败原因
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolAuthFailedMsg = sync.Pool{
	New: func() any {
		return new(AuthFailedMsg)
	},
}

// GetAuthFailedMsg() 从对象池中获取AuthFailedMsg
func GetAuthFailedMsg() *AuthFailedMsg {
	return poolAuthFailedMsg.Get().(*AuthFailedMsg)
}

// ReleaseAuthFailedMsg 释放AuthFailedMsg
func ReleaseAuthFailedMsg(v *AuthFailedMsg) {
	v.Phone = ""
	v.Msg = ""
	poolAuthFailedMsg.Put(v)
}
