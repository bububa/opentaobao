package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydListupoutResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydListupoutResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydListupoutResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydListupoutResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydListupoutResultModel
func GetAlibabaAlihealthDrugtraceTopLsydListupoutResultModel() *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydListupoutResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydListupoutResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutResultModel 释放AlibabaAlihealthDrugtraceTopLsydListupoutResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutResultModel(v *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydListupoutResultModel.Put(v)
}
