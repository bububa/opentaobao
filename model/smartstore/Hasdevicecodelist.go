package smartstore

import (
	"sync"
)

// Hasdevicecodelist 结构体
type Hasdevicecodelist struct {
	// 设备名称
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 设备类型
	DeviceType string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 设备code
	DeviceCode string `json:"device_code,omitempty" xml:"device_code,omitempty"`
}

var poolHasdevicecodelist = sync.Pool{
	New: func() any {
		return new(Hasdevicecodelist)
	},
}

// GetHasdevicecodelist() 从对象池中获取Hasdevicecodelist
func GetHasdevicecodelist() *Hasdevicecodelist {
	return poolHasdevicecodelist.Get().(*Hasdevicecodelist)
}

// ReleaseHasdevicecodelist 释放Hasdevicecodelist
func ReleaseHasdevicecodelist(v *Hasdevicecodelist) {
	v.DeviceName = ""
	v.DeviceType = ""
	v.DeviceCode = ""
	poolHasdevicecodelist.Put(v)
}
