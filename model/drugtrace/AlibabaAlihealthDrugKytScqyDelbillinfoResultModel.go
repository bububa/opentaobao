package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqyDelbillinfoResultModel 结构体
type AlibabaAlihealthDrugKytScqyDelbillinfoResultModel struct {
	// 内层对象
	Models string `json:"models,omitempty" xml:"models,omitempty"`
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqyDelbillinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyDelbillinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqyDelbillinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqyDelbillinfoResultModel
func GetAlibabaAlihealthDrugKytScqyDelbillinfoResultModel() *AlibabaAlihealthDrugKytScqyDelbillinfoResultModel {
	return poolAlibabaAlihealthDrugKytScqyDelbillinfoResultModel.Get().(*AlibabaAlihealthDrugKytScqyDelbillinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqyDelbillinfoResultModel 释放AlibabaAlihealthDrugKytScqyDelbillinfoResultModel
func ReleaseAlibabaAlihealthDrugKytScqyDelbillinfoResultModel(v *AlibabaAlihealthDrugKytScqyDelbillinfoResultModel) {
	v.Models = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytScqyDelbillinfoResultModel.Put(v)
}
