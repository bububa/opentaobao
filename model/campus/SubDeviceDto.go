package campus

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// SubDeviceDto 结构体
type SubDeviceDto struct {
	// 设备子类
	Kind string `json:"kind,omitempty" xml:"kind,omitempty"`
	// 电平量信号输入
	PortType string `json:"port_type,omitempty" xml:"port_type,omitempty"`
	// 子设备名称
	SubDeviceName string `json:"sub_device_name,omitempty" xml:"sub_device_name,omitempty"`
	// 出
	Direction string `json:"direction,omitempty" xml:"direction,omitempty"`
	// xxx
	SubDeviceId string `json:"sub_device_id,omitempty" xml:"sub_device_id,omitempty"`
	// 端口号
	Port *model.File `json:"port,omitempty" xml:"port,omitempty"`
}

var poolSubDeviceDto = sync.Pool{
	New: func() any {
		return new(SubDeviceDto)
	},
}

// GetSubDeviceDto() 从对象池中获取SubDeviceDto
func GetSubDeviceDto() *SubDeviceDto {
	return poolSubDeviceDto.Get().(*SubDeviceDto)
}

// ReleaseSubDeviceDto 释放SubDeviceDto
func ReleaseSubDeviceDto(v *SubDeviceDto) {
	v.Kind = ""
	v.PortType = ""
	v.SubDeviceName = ""
	v.Direction = ""
	v.SubDeviceId = ""
	v.Port = nil
	poolSubDeviceDto.Put(v)
}
