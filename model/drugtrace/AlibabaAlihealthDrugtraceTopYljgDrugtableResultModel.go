package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回模型
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel
func GetAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel() *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel 释放AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel(v *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopYljgDrugtableResultModel.Put(v)
}
