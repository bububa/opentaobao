package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytCodetobillResultModel 结构体
type AlibabaAlihealthDrugKytCodetobillResultModel struct {
	// 调用编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用结果
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	Model *CodeToBill `json:"model,omitempty" xml:"model,omitempty"`
	// 返回结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytCodetobillResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytCodetobillResultModel)
	},
}

// GetAlibabaAlihealthDrugKytCodetobillResultModel() 从对象池中获取AlibabaAlihealthDrugKytCodetobillResultModel
func GetAlibabaAlihealthDrugKytCodetobillResultModel() *AlibabaAlihealthDrugKytCodetobillResultModel {
	return poolAlibabaAlihealthDrugKytCodetobillResultModel.Get().(*AlibabaAlihealthDrugKytCodetobillResultModel)
}

// ReleaseAlibabaAlihealthDrugKytCodetobillResultModel 释放AlibabaAlihealthDrugKytCodetobillResultModel
func ReleaseAlibabaAlihealthDrugKytCodetobillResultModel(v *AlibabaAlihealthDrugKytCodetobillResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytCodetobillResultModel.Put(v)
}
