package campus

import (
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
