package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytListauthsResultModel 结构体
type AlibabaAlihealthDrugKytListauthsResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *Page `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytListauthsResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListauthsResultModel)
	},
}

// GetAlibabaAlihealthDrugKytListauthsResultModel() 从对象池中获取AlibabaAlihealthDrugKytListauthsResultModel
func GetAlibabaAlihealthDrugKytListauthsResultModel() *AlibabaAlihealthDrugKytListauthsResultModel {
	return poolAlibabaAlihealthDrugKytListauthsResultModel.Get().(*AlibabaAlihealthDrugKytListauthsResultModel)
}

// ReleaseAlibabaAlihealthDrugKytListauthsResultModel 释放AlibabaAlihealthDrugKytListauthsResultModel
func ReleaseAlibabaAlihealthDrugKytListauthsResultModel(v *AlibabaAlihealthDrugKytListauthsResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytListauthsResultModel.Put(v)
}
