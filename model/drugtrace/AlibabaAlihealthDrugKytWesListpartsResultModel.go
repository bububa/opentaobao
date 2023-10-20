package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesListpartsResultModel 结构体
type AlibabaAlihealthDrugKytWesListpartsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesListpartsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesListpartsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesListpartsResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesListpartsResultModel
func GetAlibabaAlihealthDrugKytWesListpartsResultModel() *AlibabaAlihealthDrugKytWesListpartsResultModel {
	return poolAlibabaAlihealthDrugKytWesListpartsResultModel.Get().(*AlibabaAlihealthDrugKytWesListpartsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesListpartsResultModel 释放AlibabaAlihealthDrugKytWesListpartsResultModel
func ReleaseAlibabaAlihealthDrugKytWesListpartsResultModel(v *AlibabaAlihealthDrugKytWesListpartsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesListpartsResultModel.Put(v)
}
