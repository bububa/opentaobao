package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel.Put(v)
}
