package iot

import (
	"sync"
)

// TopDeviceExtInfoDto 结构体
type TopDeviceExtInfoDto struct {
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 三方设备id
	ExtDeviceId string `json:"ext_device_id,omitempty" xml:"ext_device_id,omitempty"`
	// 三方设备类型
	ExtDeviceType string `json:"ext_device_type,omitempty" xml:"ext_device_type,omitempty"`
	// 设备mac
	DeviceMac string `json:"device_mac,omitempty" xml:"device_mac,omitempty"`
}

var poolTopDeviceExtInfoDto = sync.Pool{
	New: func() any {
		return new(TopDeviceExtInfoDto)
	},
}

// GetTopDeviceExtInfoDto() 从对象池中获取TopDeviceExtInfoDto
func GetTopDeviceExtInfoDto() *TopDeviceExtInfoDto {
	return poolTopDeviceExtInfoDto.Get().(*TopDeviceExtInfoDto)
}

// ReleaseTopDeviceExtInfoDto 释放TopDeviceExtInfoDto
func ReleaseTopDeviceExtInfoDto(v *TopDeviceExtInfoDto) {
	v.DeviceId = ""
	v.ExtDeviceId = ""
	v.ExtDeviceType = ""
	v.DeviceMac = ""
	poolTopDeviceExtInfoDto.Put(v)
}
