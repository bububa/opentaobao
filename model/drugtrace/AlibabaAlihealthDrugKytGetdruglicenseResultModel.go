package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetdruglicenseResultModel 结构体
type AlibabaAlihealthDrugKytGetdruglicenseResultModel struct {
	// 列表
	ModelList []AlibabaAlihealthDrugKytGetdruglicenseModel `json:"model_list,omitempty" xml:"model_list>alibaba_alihealth_drug_kyt_getdruglicense_model,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetdruglicenseResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetdruglicenseResultModel)
	},
}

// GetAlibabaAlihealthDrugKytGetdruglicenseResultModel() 从对象池中获取AlibabaAlihealthDrugKytGetdruglicenseResultModel
func GetAlibabaAlihealthDrugKytGetdruglicenseResultModel() *AlibabaAlihealthDrugKytGetdruglicenseResultModel {
	return poolAlibabaAlihealthDrugKytGetdruglicenseResultModel.Get().(*AlibabaAlihealthDrugKytGetdruglicenseResultModel)
}

// ReleaseAlibabaAlihealthDrugKytGetdruglicenseResultModel 释放AlibabaAlihealthDrugKytGetdruglicenseResultModel
func ReleaseAlibabaAlihealthDrugKytGetdruglicenseResultModel(v *AlibabaAlihealthDrugKytGetdruglicenseResultModel) {
	v.ModelList = v.ModelList[:0]
	poolAlibabaAlihealthDrugKytGetdruglicenseResultModel.Put(v)
}
