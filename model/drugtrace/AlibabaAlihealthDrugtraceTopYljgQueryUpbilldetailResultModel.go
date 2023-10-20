package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 对象模型信息
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel
func GetAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel() *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel 释放AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel(v *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel.Put(v)
}
