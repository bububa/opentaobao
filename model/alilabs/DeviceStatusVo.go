package alilabs

import (
	"sync"
)

// DeviceStatusVo 结构体
type DeviceStatusVo struct {
	// 扩展返回，保留使用
	Extensions string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// 在线状态（0：离线，1：在线）
	OnlineStatus string `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolDeviceStatusVo = sync.Pool{
	New: func() any {
		return new(DeviceStatusVo)
	},
}

// GetDeviceStatusVo() 从对象池中获取DeviceStatusVo
func GetDeviceStatusVo() *DeviceStatusVo {
	return poolDeviceStatusVo.Get().(*DeviceStatusVo)
}

// ReleaseDeviceStatusVo 释放DeviceStatusVo
func ReleaseDeviceStatusVo(v *DeviceStatusVo) {
	v.Extensions = ""
	v.OnlineStatus = ""
	v.Uuid = ""
	poolDeviceStatusVo.Put(v)
}
