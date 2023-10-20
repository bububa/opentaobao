package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel 结构体
type AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel() 从对象池中获取AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel
func GetAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel() *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel {
	return poolAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel.Get().(*AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel 释放AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel
func ReleaseAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel(v *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel.Put(v)
}
