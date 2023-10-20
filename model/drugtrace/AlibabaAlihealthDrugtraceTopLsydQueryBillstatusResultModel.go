package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel.Put(v)
}
