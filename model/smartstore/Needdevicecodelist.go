package smartstore

import (
	"sync"
)

// Needdevicecodelist 结构体
type Needdevicecodelist struct {
	// 设备名称
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 设备类型
	DeviceType string `json:"device_type,omitempty" xml:"device_type,omitempty"`
}

var poolNeeddevicecodelist = sync.Pool{
	New: func() any {
		return new(Needdevicecodelist)
	},
}

// GetNeeddevicecodelist() 从对象池中获取Needdevicecodelist
func GetNeeddevicecodelist() *Needdevicecodelist {
	return poolNeeddevicecodelist.Get().(*Needdevicecodelist)
}

// ReleaseNeeddevicecodelist 释放Needdevicecodelist
func ReleaseNeeddevicecodelist(v *Needdevicecodelist) {
	v.DeviceName = ""
	v.DeviceType = ""
	poolNeeddevicecodelist.Put(v)
}
