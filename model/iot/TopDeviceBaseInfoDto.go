package iot

import (
	"sync"
)

// TopDeviceBaseInfoDto 结构体
type TopDeviceBaseInfoDto struct {
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}

var poolTopDeviceBaseInfoDto = sync.Pool{
	New: func() any {
		return new(TopDeviceBaseInfoDto)
	},
}

// GetTopDeviceBaseInfoDto() 从对象池中获取TopDeviceBaseInfoDto
func GetTopDeviceBaseInfoDto() *TopDeviceBaseInfoDto {
	return poolTopDeviceBaseInfoDto.Get().(*TopDeviceBaseInfoDto)
}

// ReleaseTopDeviceBaseInfoDto 释放TopDeviceBaseInfoDto
func ReleaseTopDeviceBaseInfoDto(v *TopDeviceBaseInfoDto) {
	v.DeviceId = ""
	poolTopDeviceBaseInfoDto.Put(v)
}
