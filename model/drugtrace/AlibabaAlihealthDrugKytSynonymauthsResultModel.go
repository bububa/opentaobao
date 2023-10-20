package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSynonymauthsResultModel 结构体
type AlibabaAlihealthDrugKytSynonymauthsResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSynonymauthsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSynonymauthsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSynonymauthsResultModel() 从对象池中获取AlibabaAlihealthDrugKytSynonymauthsResultModel
func GetAlibabaAlihealthDrugKytSynonymauthsResultModel() *AlibabaAlihealthDrugKytSynonymauthsResultModel {
	return poolAlibabaAlihealthDrugKytSynonymauthsResultModel.Get().(*AlibabaAlihealthDrugKytSynonymauthsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSynonymauthsResultModel 释放AlibabaAlihealthDrugKytSynonymauthsResultModel
func ReleaseAlibabaAlihealthDrugKytSynonymauthsResultModel(v *AlibabaAlihealthDrugKytSynonymauthsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSynonymauthsResultModel.Put(v)
}
