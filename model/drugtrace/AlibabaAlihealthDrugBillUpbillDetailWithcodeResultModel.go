package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel 结构体
type AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel struct {
	// 提示信息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 提示信息内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 最外层对象
	Model *BillUpOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel() 从对象池中获取AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel
func GetAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel() *AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel {
	return poolAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel.Get().(*AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel 释放AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel
func ReleaseAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel(v *AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel.Put(v)
}
