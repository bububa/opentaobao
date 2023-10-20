package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugLsydSaveentResultModel 结构体
type AlibabaAlihealthDrugLsydSaveentResultModel struct {
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 具体返回值
	Model *AlibabaAlihealthDrugLsydSaveentModel `json:"model,omitempty" xml:"model,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugLsydSaveentResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugLsydSaveentResultModel)
	},
}

// GetAlibabaAlihealthDrugLsydSaveentResultModel() 从对象池中获取AlibabaAlihealthDrugLsydSaveentResultModel
func GetAlibabaAlihealthDrugLsydSaveentResultModel() *AlibabaAlihealthDrugLsydSaveentResultModel {
	return poolAlibabaAlihealthDrugLsydSaveentResultModel.Get().(*AlibabaAlihealthDrugLsydSaveentResultModel)
}

// ReleaseAlibabaAlihealthDrugLsydSaveentResultModel 释放AlibabaAlihealthDrugLsydSaveentResultModel
func ReleaseAlibabaAlihealthDrugLsydSaveentResultModel(v *AlibabaAlihealthDrugLsydSaveentResultModel) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugLsydSaveentResultModel.Put(v)
}
