package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSmyxListpartsResultModel 结构体
type AlibabaAlihealthDrugKytSmyxListpartsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSmyxListpartsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSmyxListpartsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSmyxListpartsResultModel() 从对象池中获取AlibabaAlihealthDrugKytSmyxListpartsResultModel
func GetAlibabaAlihealthDrugKytSmyxListpartsResultModel() *AlibabaAlihealthDrugKytSmyxListpartsResultModel {
	return poolAlibabaAlihealthDrugKytSmyxListpartsResultModel.Get().(*AlibabaAlihealthDrugKytSmyxListpartsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSmyxListpartsResultModel 释放AlibabaAlihealthDrugKytSmyxListpartsResultModel
func ReleaseAlibabaAlihealthDrugKytSmyxListpartsResultModel(v *AlibabaAlihealthDrugKytSmyxListpartsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSmyxListpartsResultModel.Put(v)
}
