package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytListpartsResultModel 结构体
type AlibabaAlihealthDrugKytListpartsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytListpartsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListpartsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytListpartsResultModel() 从对象池中获取AlibabaAlihealthDrugKytListpartsResultModel
func GetAlibabaAlihealthDrugKytListpartsResultModel() *AlibabaAlihealthDrugKytListpartsResultModel {
	return poolAlibabaAlihealthDrugKytListpartsResultModel.Get().(*AlibabaAlihealthDrugKytListpartsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytListpartsResultModel 释放AlibabaAlihealthDrugKytListpartsResultModel
func ReleaseAlibabaAlihealthDrugKytListpartsResultModel(v *AlibabaAlihealthDrugKytListpartsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytListpartsResultModel.Put(v)
}
