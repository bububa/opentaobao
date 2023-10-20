package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrGetbyrefentidResultModel 结构体
type AlibabaAlihealthDrugKytDrGetbyrefentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrGetbyrefentidResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetbyrefentidResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrGetbyrefentidResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrGetbyrefentidResultModel
func GetAlibabaAlihealthDrugKytDrGetbyrefentidResultModel() *AlibabaAlihealthDrugKytDrGetbyrefentidResultModel {
	return poolAlibabaAlihealthDrugKytDrGetbyrefentidResultModel.Get().(*AlibabaAlihealthDrugKytDrGetbyrefentidResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrGetbyrefentidResultModel 释放AlibabaAlihealthDrugKytDrGetbyrefentidResultModel
func ReleaseAlibabaAlihealthDrugKytDrGetbyrefentidResultModel(v *AlibabaAlihealthDrugKytDrGetbyrefentidResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrGetbyrefentidResultModel.Put(v)
}
