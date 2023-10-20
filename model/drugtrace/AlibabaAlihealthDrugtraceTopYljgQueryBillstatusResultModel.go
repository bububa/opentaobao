package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel
func GetAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel() *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel 释放AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel(v *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel.Put(v)
}
