package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqyCodeprocessResultModel 结构体
type AlibabaAlihealthDrugKytScqyCodeprocessResultModel struct {
	// 错误信息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标识
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqyCodeprocessResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyCodeprocessResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqyCodeprocessResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqyCodeprocessResultModel
func GetAlibabaAlihealthDrugKytScqyCodeprocessResultModel() *AlibabaAlihealthDrugKytScqyCodeprocessResultModel {
	return poolAlibabaAlihealthDrugKytScqyCodeprocessResultModel.Get().(*AlibabaAlihealthDrugKytScqyCodeprocessResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqyCodeprocessResultModel 释放AlibabaAlihealthDrugKytScqyCodeprocessResultModel
func ReleaseAlibabaAlihealthDrugKytScqyCodeprocessResultModel(v *AlibabaAlihealthDrugKytScqyCodeprocessResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytScqyCodeprocessResultModel.Put(v)
}
