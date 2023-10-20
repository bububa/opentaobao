package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytYyListpartsResultModel 结构体
type AlibabaAlihealthDrugKytYyListpartsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytYyListpartsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYyListpartsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytYyListpartsResultModel() 从对象池中获取AlibabaAlihealthDrugKytYyListpartsResultModel
func GetAlibabaAlihealthDrugKytYyListpartsResultModel() *AlibabaAlihealthDrugKytYyListpartsResultModel {
	return poolAlibabaAlihealthDrugKytYyListpartsResultModel.Get().(*AlibabaAlihealthDrugKytYyListpartsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytYyListpartsResultModel 释放AlibabaAlihealthDrugKytYyListpartsResultModel
func ReleaseAlibabaAlihealthDrugKytYyListpartsResultModel(v *AlibabaAlihealthDrugKytYyListpartsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytYyListpartsResultModel.Put(v)
}
