package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetproteminfoModel 结构体
type AlibabaAlihealthDrugKytDrGetproteminfoModel struct {
	// 存储温度
	StorageTemperatureList []StorageTemperatureList `json:"storage_temperature_list,omitempty" xml:"storage_temperature_list>storage_temperature_list,omitempty"`
	// 运输温度
	TransportTemperatureList []TransportTemperatureList `json:"transport_temperature_list,omitempty" xml:"transport_temperature_list>transport_temperature_list,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetproteminfoModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetproteminfoModel)
	},
}

// GetAlibabaAlihealthDrugKytDrGetproteminfoModel() 从对象池中获取AlibabaAlihealthDrugKytDrGetproteminfoModel
func GetAlibabaAlihealthDrugKytDrGetproteminfoModel() *AlibabaAlihealthDrugKytDrGetproteminfoModel {
	return poolAlibabaAlihealthDrugKytDrGetproteminfoModel.Get().(*AlibabaAlihealthDrugKytDrGetproteminfoModel)
}

// ReleaseAlibabaAlihealthDrugKytDrGetproteminfoModel 释放AlibabaAlihealthDrugKytDrGetproteminfoModel
func ReleaseAlibabaAlihealthDrugKytDrGetproteminfoModel(v *AlibabaAlihealthDrugKytDrGetproteminfoModel) {
	v.StorageTemperatureList = v.StorageTemperatureList[:0]
	v.TransportTemperatureList = v.TransportTemperatureList[:0]
	poolAlibabaAlihealthDrugKytDrGetproteminfoModel.Put(v)
}
