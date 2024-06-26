package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgListupoutResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgListupoutResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgListupoutResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgListupoutResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgListupoutResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgListupoutResultModel
func GetAlibabaAlihealthDrugtraceTopYljgListupoutResultModel() *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgListupoutResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgListupoutResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgListupoutResultModel 释放AlibabaAlihealthDrugtraceTopYljgListupoutResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgListupoutResultModel(v *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopYljgListupoutResultModel.Put(v)
}
