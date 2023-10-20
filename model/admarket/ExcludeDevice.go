package admarket

import (
	"sync"
)

// ExcludeDevice 结构体
type ExcludeDevice struct {
	// 排他设备id
	DeviceIds []string `json:"device_ids,omitempty" xml:"device_ids>string,omitempty"`
	// 排他开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 排他结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolExcludeDevice = sync.Pool{
	New: func() any {
		return new(ExcludeDevice)
	},
}

// GetExcludeDevice() 从对象池中获取ExcludeDevice
func GetExcludeDevice() *ExcludeDevice {
	return poolExcludeDevice.Get().(*ExcludeDevice)
}

// ReleaseExcludeDevice 释放ExcludeDevice
func ReleaseExcludeDevice(v *ExcludeDevice) {
	v.DeviceIds = v.DeviceIds[:0]
	v.StartTime = ""
	v.EndTime = ""
	poolExcludeDevice.Put(v)
}
