package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetentlicenseResultModel 结构体
type AlibabaAlihealthDrugKytGetentlicenseResultModel struct {
	// 列表
	ModelList []AlibabaAlihealthDrugKytGetentlicenseModel `json:"model_list,omitempty" xml:"model_list>alibaba_alihealth_drug_kyt_getentlicense_model,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetentlicenseResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetentlicenseResultModel)
	},
}

// GetAlibabaAlihealthDrugKytGetentlicenseResultModel() 从对象池中获取AlibabaAlihealthDrugKytGetentlicenseResultModel
func GetAlibabaAlihealthDrugKytGetentlicenseResultModel() *AlibabaAlihealthDrugKytGetentlicenseResultModel {
	return poolAlibabaAlihealthDrugKytGetentlicenseResultModel.Get().(*AlibabaAlihealthDrugKytGetentlicenseResultModel)
}

// ReleaseAlibabaAlihealthDrugKytGetentlicenseResultModel 释放AlibabaAlihealthDrugKytGetentlicenseResultModel
func ReleaseAlibabaAlihealthDrugKytGetentlicenseResultModel(v *AlibabaAlihealthDrugKytGetentlicenseResultModel) {
	v.ModelList = v.ModelList[:0]
	poolAlibabaAlihealthDrugKytGetentlicenseResultModel.Put(v)
}
