package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesSearchbillResultModel 结构体
type AlibabaAlihealthDrugKytWesSearchbillResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回模型
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesSearchbillResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSearchbillResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesSearchbillResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesSearchbillResultModel
func GetAlibabaAlihealthDrugKytWesSearchbillResultModel() *AlibabaAlihealthDrugKytWesSearchbillResultModel {
	return poolAlibabaAlihealthDrugKytWesSearchbillResultModel.Get().(*AlibabaAlihealthDrugKytWesSearchbillResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchbillResultModel 释放AlibabaAlihealthDrugKytWesSearchbillResultModel
func ReleaseAlibabaAlihealthDrugKytWesSearchbillResultModel(v *AlibabaAlihealthDrugKytWesSearchbillResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytWesSearchbillResultModel.Put(v)
}
