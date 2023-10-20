package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesListpartsByagentResultModel 结构体
type AlibabaAlihealthDrugKytWesListpartsByagentResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesListpartsByagentResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesListpartsByagentResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesListpartsByagentResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesListpartsByagentResultModel
func GetAlibabaAlihealthDrugKytWesListpartsByagentResultModel() *AlibabaAlihealthDrugKytWesListpartsByagentResultModel {
	return poolAlibabaAlihealthDrugKytWesListpartsByagentResultModel.Get().(*AlibabaAlihealthDrugKytWesListpartsByagentResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesListpartsByagentResultModel 释放AlibabaAlihealthDrugKytWesListpartsByagentResultModel
func ReleaseAlibabaAlihealthDrugKytWesListpartsByagentResultModel(v *AlibabaAlihealthDrugKytWesListpartsByagentResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesListpartsByagentResultModel.Put(v)
}
