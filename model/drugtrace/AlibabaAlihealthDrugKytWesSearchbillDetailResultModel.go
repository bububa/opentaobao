package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesSearchbillDetailResultModel 结构体
type AlibabaAlihealthDrugKytWesSearchbillDetailResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 对象模型信息
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesSearchbillDetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSearchbillDetailResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesSearchbillDetailResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesSearchbillDetailResultModel
func GetAlibabaAlihealthDrugKytWesSearchbillDetailResultModel() *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel {
	return poolAlibabaAlihealthDrugKytWesSearchbillDetailResultModel.Get().(*AlibabaAlihealthDrugKytWesSearchbillDetailResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchbillDetailResultModel 释放AlibabaAlihealthDrugKytWesSearchbillDetailResultModel
func ReleaseAlibabaAlihealthDrugKytWesSearchbillDetailResultModel(v *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesSearchbillDetailResultModel.Put(v)
}
