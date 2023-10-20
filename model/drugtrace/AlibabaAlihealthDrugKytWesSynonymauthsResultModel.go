package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesSynonymauthsResultModel 结构体
type AlibabaAlihealthDrugKytWesSynonymauthsResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesSynonymauthsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSynonymauthsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesSynonymauthsResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesSynonymauthsResultModel
func GetAlibabaAlihealthDrugKytWesSynonymauthsResultModel() *AlibabaAlihealthDrugKytWesSynonymauthsResultModel {
	return poolAlibabaAlihealthDrugKytWesSynonymauthsResultModel.Get().(*AlibabaAlihealthDrugKytWesSynonymauthsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesSynonymauthsResultModel 释放AlibabaAlihealthDrugKytWesSynonymauthsResultModel
func ReleaseAlibabaAlihealthDrugKytWesSynonymauthsResultModel(v *AlibabaAlihealthDrugKytWesSynonymauthsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesSynonymauthsResultModel.Put(v)
}
