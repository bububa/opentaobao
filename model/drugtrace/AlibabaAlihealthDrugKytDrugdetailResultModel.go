package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrugdetailResultModel 结构体
type AlibabaAlihealthDrugKytDrugdetailResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *ResDrugDetailInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrugdetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugdetailResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrugdetailResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrugdetailResultModel
func GetAlibabaAlihealthDrugKytDrugdetailResultModel() *AlibabaAlihealthDrugKytDrugdetailResultModel {
	return poolAlibabaAlihealthDrugKytDrugdetailResultModel.Get().(*AlibabaAlihealthDrugKytDrugdetailResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrugdetailResultModel 释放AlibabaAlihealthDrugKytDrugdetailResultModel
func ReleaseAlibabaAlihealthDrugKytDrugdetailResultModel(v *AlibabaAlihealthDrugKytDrugdetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrugdetailResultModel.Put(v)
}
