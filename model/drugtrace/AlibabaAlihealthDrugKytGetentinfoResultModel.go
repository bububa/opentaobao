package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetentinfoResultModel 结构体
type AlibabaAlihealthDrugKytGetentinfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetentinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetentinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytGetentinfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytGetentinfoResultModel
func GetAlibabaAlihealthDrugKytGetentinfoResultModel() *AlibabaAlihealthDrugKytGetentinfoResultModel {
	return poolAlibabaAlihealthDrugKytGetentinfoResultModel.Get().(*AlibabaAlihealthDrugKytGetentinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytGetentinfoResultModel 释放AlibabaAlihealthDrugKytGetentinfoResultModel
func ReleaseAlibabaAlihealthDrugKytGetentinfoResultModel(v *AlibabaAlihealthDrugKytGetentinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytGetentinfoResultModel.Put(v)
}
