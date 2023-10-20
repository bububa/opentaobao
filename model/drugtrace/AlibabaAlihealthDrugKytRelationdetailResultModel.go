package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytRelationdetailResultModel 结构体
type AlibabaAlihealthDrugKytRelationdetailResultModel struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *CodeRelationDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytRelationdetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytRelationdetailResultModel)
	},
}

// GetAlibabaAlihealthDrugKytRelationdetailResultModel() 从对象池中获取AlibabaAlihealthDrugKytRelationdetailResultModel
func GetAlibabaAlihealthDrugKytRelationdetailResultModel() *AlibabaAlihealthDrugKytRelationdetailResultModel {
	return poolAlibabaAlihealthDrugKytRelationdetailResultModel.Get().(*AlibabaAlihealthDrugKytRelationdetailResultModel)
}

// ReleaseAlibabaAlihealthDrugKytRelationdetailResultModel 释放AlibabaAlihealthDrugKytRelationdetailResultModel
func ReleaseAlibabaAlihealthDrugKytRelationdetailResultModel(v *AlibabaAlihealthDrugKytRelationdetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytRelationdetailResultModel.Put(v)
}
