package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytStorebilllistResultModel 结构体
type AlibabaAlihealthDrugKytStorebilllistResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 返回是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytStorebilllistResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytStorebilllistResultModel)
	},
}

// GetAlibabaAlihealthDrugKytStorebilllistResultModel() 从对象池中获取AlibabaAlihealthDrugKytStorebilllistResultModel
func GetAlibabaAlihealthDrugKytStorebilllistResultModel() *AlibabaAlihealthDrugKytStorebilllistResultModel {
	return poolAlibabaAlihealthDrugKytStorebilllistResultModel.Get().(*AlibabaAlihealthDrugKytStorebilllistResultModel)
}

// ReleaseAlibabaAlihealthDrugKytStorebilllistResultModel 释放AlibabaAlihealthDrugKytStorebilllistResultModel
func ReleaseAlibabaAlihealthDrugKytStorebilllistResultModel(v *AlibabaAlihealthDrugKytStorebilllistResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytStorebilllistResultModel.Put(v)
}
