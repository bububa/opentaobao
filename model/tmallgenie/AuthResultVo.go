package tmallgenie

import (
	"sync"
)

// AuthResultVo 结构体
type AuthResultVo struct {
	// 设备uuid
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 精灵用户openid
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
}

var poolAuthResultVo = sync.Pool{
	New: func() any {
		return new(AuthResultVo)
	},
}

// GetAuthResultVo() 从对象池中获取AuthResultVo
func GetAuthResultVo() *AuthResultVo {
	return poolAuthResultVo.Get().(*AuthResultVo)
}

// ReleaseAuthResultVo 释放AuthResultVo
func ReleaseAuthResultVo(v *AuthResultVo) {
	v.DeviceId = ""
	v.UserOpenId = ""
	poolAuthResultVo.Put(v)
}
