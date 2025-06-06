package tvupadmin

import (
	"sync"
)

// DeviceExtendDo 结构体
type DeviceExtendDo struct {
	// 设备型号
	DeviceModel string `json:"device_model,omitempty" xml:"device_model,omitempty"`
	// 设备内部型号
	InnerDeviceModel string `json:"inner_device_model,omitempty" xml:"inner_device_model,omitempty"`
	// 设备类型
	TerminalType string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
	// 厂商名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 牌照方
	Bcp int64 `json:"bcp,omitempty" xml:"bcp,omitempty"`
	// 厂商ID
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}

var poolDeviceExtendDo = sync.Pool{
	New: func() any {
		return new(DeviceExtendDo)
	},
}

// GetDeviceExtendDo() 从对象池中获取DeviceExtendDo
func GetDeviceExtendDo() *DeviceExtendDo {
	return poolDeviceExtendDo.Get().(*DeviceExtendDo)
}

// ReleaseDeviceExtendDo 释放DeviceExtendDo
func ReleaseDeviceExtendDo(v *DeviceExtendDo) {
	v.DeviceModel = ""
	v.InnerDeviceModel = ""
	v.TerminalType = ""
	v.BrandName = ""
	v.CreateTime = ""
	v.Bcp = 0
	v.BrandId = 0
	poolDeviceExtendDo.Put(v)
}
