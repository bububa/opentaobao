package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqySearchstatusResultModel 结构体
type AlibabaAlihealthDrugKytScqySearchstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqySearchstatusResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqySearchstatusResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqySearchstatusResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqySearchstatusResultModel
func GetAlibabaAlihealthDrugKytScqySearchstatusResultModel() *AlibabaAlihealthDrugKytScqySearchstatusResultModel {
	return poolAlibabaAlihealthDrugKytScqySearchstatusResultModel.Get().(*AlibabaAlihealthDrugKytScqySearchstatusResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqySearchstatusResultModel 释放AlibabaAlihealthDrugKytScqySearchstatusResultModel
func ReleaseAlibabaAlihealthDrugKytScqySearchstatusResultModel(v *AlibabaAlihealthDrugKytScqySearchstatusResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytScqySearchstatusResultModel.Put(v)
}
