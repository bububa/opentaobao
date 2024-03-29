package campus

import (
	"sync"
)

// DeviceDataApiDto 结构体
type DeviceDataApiDto struct {
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备code
	DeviceCode string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// 参数点code
	PropertyCode string `json:"property_code,omitempty" xml:"property_code,omitempty"`
	// 历史数据值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 值类型名称
	ValueTypeName string `json:"value_type_name,omitempty" xml:"value_type_name,omitempty"`
	// 单位编码
	UnitCode string `json:"unit_code,omitempty" xml:"unit_code,omitempty"`
	// 时间戳
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 值类型id
	ValueType int64 `json:"value_type,omitempty" xml:"value_type,omitempty"`
	// 单位id
	UnitId int64 `json:"unit_id,omitempty" xml:"unit_id,omitempty"`
}

var poolDeviceDataApiDto = sync.Pool{
	New: func() any {
		return new(DeviceDataApiDto)
	},
}

// GetDeviceDataApiDto() 从对象池中获取DeviceDataApiDto
func GetDeviceDataApiDto() *DeviceDataApiDto {
	return poolDeviceDataApiDto.Get().(*DeviceDataApiDto)
}

// ReleaseDeviceDataApiDto 释放DeviceDataApiDto
func ReleaseDeviceDataApiDto(v *DeviceDataApiDto) {
	v.DeviceId = ""
	v.DeviceCode = ""
	v.PropertyCode = ""
	v.Value = ""
	v.ValueTypeName = ""
	v.UnitCode = ""
	v.Timestamp = 0
	v.ValueType = 0
	v.UnitId = 0
	poolDeviceDataApiDto.Put(v)
}
