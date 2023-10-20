package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSmyxGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytSmyxGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSmyxGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSmyxGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSmyxGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytSmyxGetentinfoResultModel
func GetAlibabaAlihealthDrugKytSmyxGetentinfoResultModel() *AlibabaAlihealthDrugKytSmyxGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytSmyxGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytSmyxGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSmyxGetentinfoResultModel 释放AlibabaAlihealthDrugKytSmyxGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytSmyxGetentinfoResultModel(v *AlibabaAlihealthDrugKytSmyxGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSmyxGetentinfoResultModel.Put(v)
}
