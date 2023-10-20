package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel
func GetAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel() *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel 释放AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel(v *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel.Put(v)
}
