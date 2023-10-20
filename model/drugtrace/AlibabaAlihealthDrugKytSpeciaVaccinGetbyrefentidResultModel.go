package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel 结构体
type AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel() 从对象池中获取AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel
func GetAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel() *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel {
	return poolAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel.Get().(*AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel 释放AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel
func ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel(v *AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidResultModel.Put(v)
}
