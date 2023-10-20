package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSearchstatusResultModel 结构体
type AlibabaAlihealthDrugKytSearchstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSearchstatusResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSearchstatusResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSearchstatusResultModel() 从对象池中获取AlibabaAlihealthDrugKytSearchstatusResultModel
func GetAlibabaAlihealthDrugKytSearchstatusResultModel() *AlibabaAlihealthDrugKytSearchstatusResultModel {
	return poolAlibabaAlihealthDrugKytSearchstatusResultModel.Get().(*AlibabaAlihealthDrugKytSearchstatusResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSearchstatusResultModel 释放AlibabaAlihealthDrugKytSearchstatusResultModel
func ReleaseAlibabaAlihealthDrugKytSearchstatusResultModel(v *AlibabaAlihealthDrugKytSearchstatusResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSearchstatusResultModel.Put(v)
}
