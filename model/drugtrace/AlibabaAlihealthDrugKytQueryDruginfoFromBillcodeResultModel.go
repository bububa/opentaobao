package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel 结构体
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel struct {
	// 消息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 模型
	Model *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 消息成功失败标记
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel
func GetAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel() *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel {
	return poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel.Get().(*AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel 释放AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel
func ReleaseAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel(v *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel.Put(v)
}
