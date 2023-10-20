package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel 结构体
type AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 内层大对象
	Model *CodeFullInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 查询成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel
func GetAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel() *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel {
	return poolAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel.Get().(*AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel 释放AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel
func ReleaseAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel(v *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel.Put(v)
}
