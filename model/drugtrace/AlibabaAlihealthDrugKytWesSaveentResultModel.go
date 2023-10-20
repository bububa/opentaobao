package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesSaveentResultModel 结构体
type AlibabaAlihealthDrugKytWesSaveentResultModel struct {
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 具体返回值
	Model *AlibabaAlihealthDrugKytWesSaveentModel `json:"model,omitempty" xml:"model,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesSaveentResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSaveentResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesSaveentResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesSaveentResultModel
func GetAlibabaAlihealthDrugKytWesSaveentResultModel() *AlibabaAlihealthDrugKytWesSaveentResultModel {
	return poolAlibabaAlihealthDrugKytWesSaveentResultModel.Get().(*AlibabaAlihealthDrugKytWesSaveentResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesSaveentResultModel 释放AlibabaAlihealthDrugKytWesSaveentResultModel
func ReleaseAlibabaAlihealthDrugKytWesSaveentResultModel(v *AlibabaAlihealthDrugKytWesSaveentResultModel) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytWesSaveentResultModel.Put(v)
}
