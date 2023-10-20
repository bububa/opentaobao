package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 对象模型信息
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel.Put(v)
}
