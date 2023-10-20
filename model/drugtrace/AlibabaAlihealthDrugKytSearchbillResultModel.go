package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSearchbillResultModel 结构体
type AlibabaAlihealthDrugKytSearchbillResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回模型
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSearchbillResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSearchbillResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSearchbillResultModel() 从对象池中获取AlibabaAlihealthDrugKytSearchbillResultModel
func GetAlibabaAlihealthDrugKytSearchbillResultModel() *AlibabaAlihealthDrugKytSearchbillResultModel {
	return poolAlibabaAlihealthDrugKytSearchbillResultModel.Get().(*AlibabaAlihealthDrugKytSearchbillResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSearchbillResultModel 释放AlibabaAlihealthDrugKytSearchbillResultModel
func ReleaseAlibabaAlihealthDrugKytSearchbillResultModel(v *AlibabaAlihealthDrugKytSearchbillResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytSearchbillResultModel.Put(v)
}
