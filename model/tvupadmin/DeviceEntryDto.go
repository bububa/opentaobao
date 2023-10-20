package tvupadmin

import (
	"sync"
)

// DeviceEntryDto 结构体
type DeviceEntryDto struct {
	// id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolDeviceEntryDto = sync.Pool{
	New: func() any {
		return new(DeviceEntryDto)
	},
}

// GetDeviceEntryDto() 从对象池中获取DeviceEntryDto
func GetDeviceEntryDto() *DeviceEntryDto {
	return poolDeviceEntryDto.Get().(*DeviceEntryDto)
}

// ReleaseDeviceEntryDto 释放DeviceEntryDto
func ReleaseDeviceEntryDto(v *DeviceEntryDto) {
	v.Id = ""
	v.Name = ""
	poolDeviceEntryDto.Put(v)
}
