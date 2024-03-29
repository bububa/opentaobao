package drugtrace

import (
	"sync"
)

// StorageTemperatureList 结构体
type StorageTemperatureList struct {
	// 设备名称
	EquipmentName string `json:"equipment_name,omitempty" xml:"equipment_name,omitempty"`
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 设备编号
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 最低温度
	MinValue string `json:"min_value,omitempty" xml:"min_value,omitempty"`
	// 上传企业名称
	UploadEntName string `json:"upload_ent_name,omitempty" xml:"upload_ent_name,omitempty"`
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 温度类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 最高温度
	MaxValue string `json:"max_value,omitempty" xml:"max_value,omitempty"`
}

var poolStorageTemperatureList = sync.Pool{
	New: func() any {
		return new(StorageTemperatureList)
	},
}

// GetStorageTemperatureList() 从对象池中获取StorageTemperatureList
func GetStorageTemperatureList() *StorageTemperatureList {
	return poolStorageTemperatureList.Get().(*StorageTemperatureList)
}

// ReleaseStorageTemperatureList 释放StorageTemperatureList
func ReleaseStorageTemperatureList(v *StorageTemperatureList) {
	v.EquipmentName = ""
	v.Time = ""
	v.EquipmentCode = ""
	v.MinValue = ""
	v.UploadEntName = ""
	v.BillCode = ""
	v.Type = ""
	v.MaxValue = ""
	poolStorageTemperatureList.Put(v)
}
