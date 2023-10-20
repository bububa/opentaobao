package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytListupoutResultModel 结构体
type AlibabaAlihealthDrugKytListupoutResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytListupoutResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListupoutResultModel)
	},
}

// GetAlibabaAlihealthDrugKytListupoutResultModel() 从对象池中获取AlibabaAlihealthDrugKytListupoutResultModel
func GetAlibabaAlihealthDrugKytListupoutResultModel() *AlibabaAlihealthDrugKytListupoutResultModel {
	return poolAlibabaAlihealthDrugKytListupoutResultModel.Get().(*AlibabaAlihealthDrugKytListupoutResultModel)
}

// ReleaseAlibabaAlihealthDrugKytListupoutResultModel 释放AlibabaAlihealthDrugKytListupoutResultModel
func ReleaseAlibabaAlihealthDrugKytListupoutResultModel(v *AlibabaAlihealthDrugKytListupoutResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytListupoutResultModel.Put(v)
}
