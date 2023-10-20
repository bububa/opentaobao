package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrugtableResultModel 结构体
type AlibabaAlihealthDrugKytDrugtableResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回模型
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrugtableResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugtableResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrugtableResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrugtableResultModel
func GetAlibabaAlihealthDrugKytDrugtableResultModel() *AlibabaAlihealthDrugKytDrugtableResultModel {
	return poolAlibabaAlihealthDrugKytDrugtableResultModel.Get().(*AlibabaAlihealthDrugKytDrugtableResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrugtableResultModel 释放AlibabaAlihealthDrugKytDrugtableResultModel
func ReleaseAlibabaAlihealthDrugKytDrugtableResultModel(v *AlibabaAlihealthDrugKytDrugtableResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrugtableResultModel.Put(v)
}
