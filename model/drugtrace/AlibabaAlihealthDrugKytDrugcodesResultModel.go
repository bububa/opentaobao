package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrugcodesResultModel 结构体
type AlibabaAlihealthDrugKytDrugcodesResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *DrugCode `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrugcodesResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugcodesResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrugcodesResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrugcodesResultModel
func GetAlibabaAlihealthDrugKytDrugcodesResultModel() *AlibabaAlihealthDrugKytDrugcodesResultModel {
	return poolAlibabaAlihealthDrugKytDrugcodesResultModel.Get().(*AlibabaAlihealthDrugKytDrugcodesResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrugcodesResultModel 释放AlibabaAlihealthDrugKytDrugcodesResultModel
func ReleaseAlibabaAlihealthDrugKytDrugcodesResultModel(v *AlibabaAlihealthDrugKytDrugcodesResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrugcodesResultModel.Put(v)
}
