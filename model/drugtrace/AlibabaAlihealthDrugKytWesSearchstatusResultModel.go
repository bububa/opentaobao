package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesSearchstatusResultModel 结构体
type AlibabaAlihealthDrugKytWesSearchstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesSearchstatusResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSearchstatusResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesSearchstatusResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesSearchstatusResultModel
func GetAlibabaAlihealthDrugKytWesSearchstatusResultModel() *AlibabaAlihealthDrugKytWesSearchstatusResultModel {
	return poolAlibabaAlihealthDrugKytWesSearchstatusResultModel.Get().(*AlibabaAlihealthDrugKytWesSearchstatusResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchstatusResultModel 释放AlibabaAlihealthDrugKytWesSearchstatusResultModel
func ReleaseAlibabaAlihealthDrugKytWesSearchstatusResultModel(v *AlibabaAlihealthDrugKytWesSearchstatusResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytWesSearchstatusResultModel.Put(v)
}
