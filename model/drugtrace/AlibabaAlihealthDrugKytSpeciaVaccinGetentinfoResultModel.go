package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel
func GetAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel() *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel 释放AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel(v *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoResultModel.Put(v)
}
