package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytYyDrugcodesResultModel 结构体
type AlibabaAlihealthDrugKytYyDrugcodesResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *DrugCode `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytYyDrugcodesResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYyDrugcodesResultModel)
	},
}

// GetAlibabaAlihealthDrugKytYyDrugcodesResultModel() 从对象池中获取AlibabaAlihealthDrugKytYyDrugcodesResultModel
func GetAlibabaAlihealthDrugKytYyDrugcodesResultModel() *AlibabaAlihealthDrugKytYyDrugcodesResultModel {
	return poolAlibabaAlihealthDrugKytYyDrugcodesResultModel.Get().(*AlibabaAlihealthDrugKytYyDrugcodesResultModel)
}

// ReleaseAlibabaAlihealthDrugKytYyDrugcodesResultModel 释放AlibabaAlihealthDrugKytYyDrugcodesResultModel
func ReleaseAlibabaAlihealthDrugKytYyDrugcodesResultModel(v *AlibabaAlihealthDrugKytYyDrugcodesResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytYyDrugcodesResultModel.Put(v)
}
