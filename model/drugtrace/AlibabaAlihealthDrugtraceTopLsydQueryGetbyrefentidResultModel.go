package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel.Put(v)
}
