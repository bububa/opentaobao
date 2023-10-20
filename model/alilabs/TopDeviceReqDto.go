package alilabs

import (
	"sync"
)

// TopDeviceReqDto 结构体
type TopDeviceReqDto struct {
	// 设备签名
	DeviceSignature string `json:"device_signature,omitempty" xml:"device_signature,omitempty"`
	// 三方设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}

var poolTopDeviceReqDto = sync.Pool{
	New: func() any {
		return new(TopDeviceReqDto)
	},
}

// GetTopDeviceReqDto() 从对象池中获取TopDeviceReqDto
func GetTopDeviceReqDto() *TopDeviceReqDto {
	return poolTopDeviceReqDto.Get().(*TopDeviceReqDto)
}

// ReleaseTopDeviceReqDto 释放TopDeviceReqDto
func ReleaseTopDeviceReqDto(v *TopDeviceReqDto) {
	v.DeviceSignature = ""
	v.DeviceId = ""
	poolTopDeviceReqDto.Put(v)
}
