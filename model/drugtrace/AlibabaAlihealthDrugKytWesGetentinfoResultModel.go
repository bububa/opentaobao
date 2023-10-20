package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytWesGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesGetentinfoResultModel
func GetAlibabaAlihealthDrugKytWesGetentinfoResultModel() *AlibabaAlihealthDrugKytWesGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytWesGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytWesGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesGetentinfoResultModel 释放AlibabaAlihealthDrugKytWesGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytWesGetentinfoResultModel(v *AlibabaAlihealthDrugKytWesGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesGetentinfoResultModel.Put(v)
}
