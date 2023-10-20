package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesGetbyrefentidResultModel 结构体
type AlibabaAlihealthDrugKytWesGetbyrefentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesGetbyrefentidResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesGetbyrefentidResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesGetbyrefentidResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesGetbyrefentidResultModel
func GetAlibabaAlihealthDrugKytWesGetbyrefentidResultModel() *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel {
	return poolAlibabaAlihealthDrugKytWesGetbyrefentidResultModel.Get().(*AlibabaAlihealthDrugKytWesGetbyrefentidResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesGetbyrefentidResultModel 释放AlibabaAlihealthDrugKytWesGetbyrefentidResultModel
func ReleaseAlibabaAlihealthDrugKytWesGetbyrefentidResultModel(v *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesGetbyrefentidResultModel.Put(v)
}
