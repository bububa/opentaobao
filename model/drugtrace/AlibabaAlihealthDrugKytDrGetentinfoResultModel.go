package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytDrGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrGetentinfoResultModel
func GetAlibabaAlihealthDrugKytDrGetentinfoResultModel() *AlibabaAlihealthDrugKytDrGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytDrGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytDrGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrGetentinfoResultModel 释放AlibabaAlihealthDrugKytDrGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytDrGetentinfoResultModel(v *AlibabaAlihealthDrugKytDrGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrGetentinfoResultModel.Put(v)
}
