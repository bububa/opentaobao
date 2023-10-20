package alilabs

import (
	"sync"
)

// DeviceTokenVo 结构体
type DeviceTokenVo struct {
	// accessToken
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// refreshToken
	RefreshToken string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// state
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// accessToken超时时间，单位为秒
	AccessExpiresIn int64 `json:"access_expires_in,omitempty" xml:"access_expires_in,omitempty"`
	// refreshToken超时时间，单位为秒
	RefreshExpiresIn int64 `json:"refresh_expires_in,omitempty" xml:"refresh_expires_in,omitempty"`
}

var poolDeviceTokenVo = sync.Pool{
	New: func() any {
		return new(DeviceTokenVo)
	},
}

// GetDeviceTokenVo() 从对象池中获取DeviceTokenVo
func GetDeviceTokenVo() *DeviceTokenVo {
	return poolDeviceTokenVo.Get().(*DeviceTokenVo)
}

// ReleaseDeviceTokenVo 释放DeviceTokenVo
func ReleaseDeviceTokenVo(v *DeviceTokenVo) {
	v.AccessToken = ""
	v.RefreshToken = ""
	v.State = ""
	v.AccessExpiresIn = 0
	v.RefreshExpiresIn = 0
	poolDeviceTokenVo.Put(v)
}
