package alilabs

import (
	"sync"
)

// DeviceInfo 结构体
type DeviceInfo struct {
	// 设备唯一id
	DevId string `json:"dev_id,omitempty" xml:"dev_id,omitempty"`
	// 设备状态Map，key和value均为String
	Status *Status `json:"status,omitempty" xml:"status,omitempty"`
}

var poolDeviceInfo = sync.Pool{
	New: func() any {
		return new(DeviceInfo)
	},
}

// GetDeviceInfo() 从对象池中获取DeviceInfo
func GetDeviceInfo() *DeviceInfo {
	return poolDeviceInfo.Get().(*DeviceInfo)
}

// ReleaseDeviceInfo 释放DeviceInfo
func ReleaseDeviceInfo(v *DeviceInfo) {
	v.DevId = ""
	v.Status = nil
	poolDeviceInfo.Put(v)
}
