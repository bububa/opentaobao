package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel 结构体
type AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel struct {
	// 成功失败编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功失败描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 成功失败结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel() 从对象池中获取AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel
func GetAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel() *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel {
	return poolAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel.Get().(*AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel 释放AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel
func ReleaseAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel(v *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel.Put(v)
}
