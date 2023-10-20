package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytVaGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytVaGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytVaGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytVaGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytVaGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytVaGetentinfoResultModel
func GetAlibabaAlihealthDrugKytVaGetentinfoResultModel() *AlibabaAlihealthDrugKytVaGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytVaGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytVaGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytVaGetentinfoResultModel 释放AlibabaAlihealthDrugKytVaGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytVaGetentinfoResultModel(v *AlibabaAlihealthDrugKytVaGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytVaGetentinfoResultModel.Put(v)
}
