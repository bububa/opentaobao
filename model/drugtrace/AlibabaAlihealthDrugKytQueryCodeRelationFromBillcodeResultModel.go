package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel 结构体
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 对象模型信息
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel
func GetAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel() *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel {
	return poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel.Get().(*AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel 释放AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel
func ReleaseAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel(v *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeResultModel.Put(v)
}
