package campus

import (
	"sync"
)

// DeviceReportDataDto 结构体
type DeviceReportDataDto struct {
	// 温度
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述信息
	Describe string `json:"describe,omitempty" xml:"describe,omitempty"`
	// 值
	OriginValue string `json:"origin_value,omitempty" xml:"origin_value,omitempty"`
	// 是否告警
	IsAlarm bool `json:"is_alarm,omitempty" xml:"is_alarm,omitempty"`
}

var poolDeviceReportDataDto = sync.Pool{
	New: func() any {
		return new(DeviceReportDataDto)
	},
}

// GetDeviceReportDataDto() 从对象池中获取DeviceReportDataDto
func GetDeviceReportDataDto() *DeviceReportDataDto {
	return poolDeviceReportDataDto.Get().(*DeviceReportDataDto)
}

// ReleaseDeviceReportDataDto 释放DeviceReportDataDto
func ReleaseDeviceReportDataDto(v *DeviceReportDataDto) {
	v.Code = ""
	v.Describe = ""
	v.OriginValue = ""
	v.IsAlarm = false
	poolDeviceReportDataDto.Put(v)
}
