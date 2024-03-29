package drugtrace

import (
	"sync"
)

// TransportTemperatureList 结构体
type TransportTemperatureList struct {
	// 设备名称
	EquipmentName string `json:"equipment_name,omitempty" xml:"equipment_name,omitempty"`
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 设备编号
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 最低温度
	MinValue string `json:"min_value,omitempty" xml:"min_value,omitempty"`
	// 上传企业
	UploadEntName string `json:"upload_ent_name,omitempty" xml:"upload_ent_name,omitempty"`
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 温度类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 最高温度
	MaxValue string `json:"max_value,omitempty" xml:"max_value,omitempty"`
}

var poolTransportTemperatureList = sync.Pool{
	New: func() any {
		return new(TransportTemperatureList)
	},
}

// GetTransportTemperatureList() 从对象池中获取TransportTemperatureList
func GetTransportTemperatureList() *TransportTemperatureList {
	return poolTransportTemperatureList.Get().(*TransportTemperatureList)
}

// ReleaseTransportTemperatureList 释放TransportTemperatureList
func ReleaseTransportTemperatureList(v *TransportTemperatureList) {
	v.EquipmentName = ""
	v.Time = ""
	v.EquipmentCode = ""
	v.MinValue = ""
	v.UploadEntName = ""
	v.BillCode = ""
	v.Type = ""
	v.MaxValue = ""
	poolTransportTemperatureList.Put(v)
}
