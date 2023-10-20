package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytCodeprocessResultModel 结构体
type AlibabaAlihealthDrugKytCodeprocessResultModel struct {
	// 错误信息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标识
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytCodeprocessResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytCodeprocessResultModel)
	},
}

// GetAlibabaAlihealthDrugKytCodeprocessResultModel() 从对象池中获取AlibabaAlihealthDrugKytCodeprocessResultModel
func GetAlibabaAlihealthDrugKytCodeprocessResultModel() *AlibabaAlihealthDrugKytCodeprocessResultModel {
	return poolAlibabaAlihealthDrugKytCodeprocessResultModel.Get().(*AlibabaAlihealthDrugKytCodeprocessResultModel)
}

// ReleaseAlibabaAlihealthDrugKytCodeprocessResultModel 释放AlibabaAlihealthDrugKytCodeprocessResultModel
func ReleaseAlibabaAlihealthDrugKytCodeprocessResultModel(v *AlibabaAlihealthDrugKytCodeprocessResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytCodeprocessResultModel.Put(v)
}
