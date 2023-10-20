package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesDrugrescodeResultModel 结构体
type AlibabaAlihealthDrugKytWesDrugrescodeResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回模型
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesDrugrescodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesDrugrescodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesDrugrescodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesDrugrescodeResultModel
func GetAlibabaAlihealthDrugKytWesDrugrescodeResultModel() *AlibabaAlihealthDrugKytWesDrugrescodeResultModel {
	return poolAlibabaAlihealthDrugKytWesDrugrescodeResultModel.Get().(*AlibabaAlihealthDrugKytWesDrugrescodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesDrugrescodeResultModel 释放AlibabaAlihealthDrugKytWesDrugrescodeResultModel
func ReleaseAlibabaAlihealthDrugKytWesDrugrescodeResultModel(v *AlibabaAlihealthDrugKytWesDrugrescodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesDrugrescodeResultModel.Put(v)
}
