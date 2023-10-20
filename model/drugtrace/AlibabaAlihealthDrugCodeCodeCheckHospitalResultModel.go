package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel 结构体
type AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel struct {
	// 成功失败编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功失败描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 成功失败结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel() 从对象池中获取AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel
func GetAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel() *AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel {
	return poolAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel.Get().(*AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel 释放AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel
func ReleaseAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel(v *AlibabaAlihealthDrugCodeCodeCheckHospitalResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthDrugCodeCodeCheckHospitalResultModel.Put(v)
}
