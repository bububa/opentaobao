package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel 结构体
type AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel struct {
	// 消息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 模型
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 消息成功失败标记
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel
func GetAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel() *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel {
	return poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel.Get().(*AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel 释放AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel
func ReleaseAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel(v *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeResultModel.Put(v)
}
