package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrSearchstatusResultModel 结构体
type AlibabaAlihealthDrugKytDrSearchstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrSearchstatusResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrSearchstatusResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrSearchstatusResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrSearchstatusResultModel
func GetAlibabaAlihealthDrugKytDrSearchstatusResultModel() *AlibabaAlihealthDrugKytDrSearchstatusResultModel {
	return poolAlibabaAlihealthDrugKytDrSearchstatusResultModel.Get().(*AlibabaAlihealthDrugKytDrSearchstatusResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrSearchstatusResultModel 释放AlibabaAlihealthDrugKytDrSearchstatusResultModel
func ReleaseAlibabaAlihealthDrugKytDrSearchstatusResultModel(v *AlibabaAlihealthDrugKytDrSearchstatusResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrSearchstatusResultModel.Put(v)
}
