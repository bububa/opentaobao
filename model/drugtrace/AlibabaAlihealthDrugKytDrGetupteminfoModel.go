package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetupteminfoModel 结构体
type AlibabaAlihealthDrugKytDrGetupteminfoModel struct {
	// 返回结果列表
	List []VaTemperatureBillResultDto `json:"list,omitempty" xml:"list>va_temperature_bill_result_dto,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetupteminfoModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetupteminfoModel)
	},
}

// GetAlibabaAlihealthDrugKytDrGetupteminfoModel() 从对象池中获取AlibabaAlihealthDrugKytDrGetupteminfoModel
func GetAlibabaAlihealthDrugKytDrGetupteminfoModel() *AlibabaAlihealthDrugKytDrGetupteminfoModel {
	return poolAlibabaAlihealthDrugKytDrGetupteminfoModel.Get().(*AlibabaAlihealthDrugKytDrGetupteminfoModel)
}

// ReleaseAlibabaAlihealthDrugKytDrGetupteminfoModel 释放AlibabaAlihealthDrugKytDrGetupteminfoModel
func ReleaseAlibabaAlihealthDrugKytDrGetupteminfoModel(v *AlibabaAlihealthDrugKytDrGetupteminfoModel) {
	v.List = v.List[:0]
	poolAlibabaAlihealthDrugKytDrGetupteminfoModel.Put(v)
}
