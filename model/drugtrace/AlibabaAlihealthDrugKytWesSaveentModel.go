package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesSaveentModel 结构体
type AlibabaAlihealthDrugKytWesSaveentModel struct {
	// 新增失败的时候错误原因
	CheckMsg string `json:"check_msg,omitempty" xml:"check_msg,omitempty"`
	// 新增成功后分配的往来单位refEntId
	ParRefEntId string `json:"par_ref_ent_id,omitempty" xml:"par_ref_ent_id,omitempty"`
	// 新增成功还是失败，true：新增成功
	AddSucess bool `json:"add_sucess,omitempty" xml:"add_sucess,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesSaveentModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSaveentModel)
	},
}

// GetAlibabaAlihealthDrugKytWesSaveentModel() 从对象池中获取AlibabaAlihealthDrugKytWesSaveentModel
func GetAlibabaAlihealthDrugKytWesSaveentModel() *AlibabaAlihealthDrugKytWesSaveentModel {
	return poolAlibabaAlihealthDrugKytWesSaveentModel.Get().(*AlibabaAlihealthDrugKytWesSaveentModel)
}

// ReleaseAlibabaAlihealthDrugKytWesSaveentModel 释放AlibabaAlihealthDrugKytWesSaveentModel
func ReleaseAlibabaAlihealthDrugKytWesSaveentModel(v *AlibabaAlihealthDrugKytWesSaveentModel) {
	v.CheckMsg = ""
	v.ParRefEntId = ""
	v.AddSucess = false
	poolAlibabaAlihealthDrugKytWesSaveentModel.Put(v)
}
