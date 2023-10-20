package alilabs

import (
	"sync"
)

// DeviceStatusDto 结构体
type DeviceStatusDto struct {
	// payload
	Payload *Payload `json:"payload,omitempty" xml:"payload,omitempty"`
	// header
	Header *IotCommonHeader `json:"header,omitempty" xml:"header,omitempty"`
}

var poolDeviceStatusDto = sync.Pool{
	New: func() any {
		return new(DeviceStatusDto)
	},
}

// GetDeviceStatusDto() 从对象池中获取DeviceStatusDto
func GetDeviceStatusDto() *DeviceStatusDto {
	return poolDeviceStatusDto.Get().(*DeviceStatusDto)
}

// ReleaseDeviceStatusDto 释放DeviceStatusDto
func ReleaseDeviceStatusDto(v *DeviceStatusDto) {
	v.Payload = nil
	v.Header = nil
	poolDeviceStatusDto.Put(v)
}
