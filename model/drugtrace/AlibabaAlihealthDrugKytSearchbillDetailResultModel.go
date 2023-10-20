package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSearchbillDetailResultModel 结构体
type AlibabaAlihealthDrugKytSearchbillDetailResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 对象模型信息
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSearchbillDetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSearchbillDetailResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSearchbillDetailResultModel() 从对象池中获取AlibabaAlihealthDrugKytSearchbillDetailResultModel
func GetAlibabaAlihealthDrugKytSearchbillDetailResultModel() *AlibabaAlihealthDrugKytSearchbillDetailResultModel {
	return poolAlibabaAlihealthDrugKytSearchbillDetailResultModel.Get().(*AlibabaAlihealthDrugKytSearchbillDetailResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSearchbillDetailResultModel 释放AlibabaAlihealthDrugKytSearchbillDetailResultModel
func ReleaseAlibabaAlihealthDrugKytSearchbillDetailResultModel(v *AlibabaAlihealthDrugKytSearchbillDetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSearchbillDetailResultModel.Put(v)
}
