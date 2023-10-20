package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel 结构体
type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel struct {
	// 提示信息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 提示信息内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 最外层对象
	Model *BillUpOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel
func GetAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel() *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel {
	return poolAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel.Get().(*AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel 释放AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel
func ReleaseAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel(v *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel.Put(v)
}
