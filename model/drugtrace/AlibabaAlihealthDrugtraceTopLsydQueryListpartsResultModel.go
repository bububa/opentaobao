package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel.Put(v)
}
