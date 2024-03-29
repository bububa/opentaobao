package campus

import (
	"sync"
)

// DeviceReportEventDto 结构体
type DeviceReportEventDto struct {
	// 上传数据
	Data []DeviceReportDataDto `json:"data,omitempty" xml:"data>device_report_data_dto,omitempty"`
	// 消息唯一id
	TransId string `json:"trans_id,omitempty" xml:"trans_id,omitempty"`
	// 应用key
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 来源系统
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 消息版本
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 0:设备中心UUID ，1：外部id
	IdType int64 `json:"id_type,omitempty" xml:"id_type,omitempty"`
	// 消息时间戳
	EventTime int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
}

var poolDeviceReportEventDto = sync.Pool{
	New: func() any {
		return new(DeviceReportEventDto)
	},
}

// GetDeviceReportEventDto() 从对象池中获取DeviceReportEventDto
func GetDeviceReportEventDto() *DeviceReportEventDto {
	return poolDeviceReportEventDto.Get().(*DeviceReportEventDto)
}

// ReleaseDeviceReportEventDto 释放DeviceReportEventDto
func ReleaseDeviceReportEventDto(v *DeviceReportEventDto) {
	v.Data = v.Data[:0]
	v.TransId = ""
	v.AppKey = ""
	v.Source = ""
	v.Version = ""
	v.DeviceId = ""
	v.IdType = 0
	v.EventTime = 0
	poolDeviceReportEventDto.Put(v)
}
