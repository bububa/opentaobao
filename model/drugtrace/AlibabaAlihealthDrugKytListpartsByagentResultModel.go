package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytListpartsByagentResultModel 结构体
type AlibabaAlihealthDrugKytListpartsByagentResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytListpartsByagentResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListpartsByagentResultModel)
	},
}

// GetAlibabaAlihealthDrugKytListpartsByagentResultModel() 从对象池中获取AlibabaAlihealthDrugKytListpartsByagentResultModel
func GetAlibabaAlihealthDrugKytListpartsByagentResultModel() *AlibabaAlihealthDrugKytListpartsByagentResultModel {
	return poolAlibabaAlihealthDrugKytListpartsByagentResultModel.Get().(*AlibabaAlihealthDrugKytListpartsByagentResultModel)
}

// ReleaseAlibabaAlihealthDrugKytListpartsByagentResultModel 释放AlibabaAlihealthDrugKytListpartsByagentResultModel
func ReleaseAlibabaAlihealthDrugKytListpartsByagentResultModel(v *AlibabaAlihealthDrugKytListpartsByagentResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytListpartsByagentResultModel.Put(v)
}
