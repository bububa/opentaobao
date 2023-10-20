package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetbyentidResultModel 结构体
type AlibabaAlihealthDrugKytDrGetbyentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetbyentidResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetbyentidResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrGetbyentidResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrGetbyentidResultModel
func GetAlibabaAlihealthDrugKytDrGetbyentidResultModel() *AlibabaAlihealthDrugKytDrGetbyentidResultModel {
	return poolAlibabaAlihealthDrugKytDrGetbyentidResultModel.Get().(*AlibabaAlihealthDrugKytDrGetbyentidResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrGetbyentidResultModel 释放AlibabaAlihealthDrugKytDrGetbyentidResultModel
func ReleaseAlibabaAlihealthDrugKytDrGetbyentidResultModel(v *AlibabaAlihealthDrugKytDrGetbyentidResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrGetbyentidResultModel.Put(v)
}
