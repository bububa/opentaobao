package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetbyentidResultModel 结构体
type AlibabaAlihealthDrugKytGetbyentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetbyentidResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetbyentidResultModel)
	},
}

// GetAlibabaAlihealthDrugKytGetbyentidResultModel() 从对象池中获取AlibabaAlihealthDrugKytGetbyentidResultModel
func GetAlibabaAlihealthDrugKytGetbyentidResultModel() *AlibabaAlihealthDrugKytGetbyentidResultModel {
	return poolAlibabaAlihealthDrugKytGetbyentidResultModel.Get().(*AlibabaAlihealthDrugKytGetbyentidResultModel)
}

// ReleaseAlibabaAlihealthDrugKytGetbyentidResultModel 释放AlibabaAlihealthDrugKytGetbyentidResultModel
func ReleaseAlibabaAlihealthDrugKytGetbyentidResultModel(v *AlibabaAlihealthDrugKytGetbyentidResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytGetbyentidResultModel.Put(v)
}
