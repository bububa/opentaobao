package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytYyGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytYyGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytYyGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYyGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytYyGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytYyGetentinfoResultModel
func GetAlibabaAlihealthDrugKytYyGetentinfoResultModel() *AlibabaAlihealthDrugKytYyGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytYyGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytYyGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytYyGetentinfoResultModel 释放AlibabaAlihealthDrugKytYyGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytYyGetentinfoResultModel(v *AlibabaAlihealthDrugKytYyGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytYyGetentinfoResultModel.Put(v)
}
