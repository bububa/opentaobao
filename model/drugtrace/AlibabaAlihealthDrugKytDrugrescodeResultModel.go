package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrugrescodeResultModel 结构体
type AlibabaAlihealthDrugKytDrugrescodeResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回模型
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrugrescodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugrescodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrugrescodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrugrescodeResultModel
func GetAlibabaAlihealthDrugKytDrugrescodeResultModel() *AlibabaAlihealthDrugKytDrugrescodeResultModel {
	return poolAlibabaAlihealthDrugKytDrugrescodeResultModel.Get().(*AlibabaAlihealthDrugKytDrugrescodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrugrescodeResultModel 释放AlibabaAlihealthDrugKytDrugrescodeResultModel
func ReleaseAlibabaAlihealthDrugKytDrugrescodeResultModel(v *AlibabaAlihealthDrugKytDrugrescodeResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrugrescodeResultModel.Put(v)
}
