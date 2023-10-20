package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrSinglerelationResultModel 结构体
type AlibabaAlihealthDrugKytDrSinglerelationResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrSinglerelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrSinglerelationResultModel)
	},
}

// GetAlibabaAlihealthDrugKytDrSinglerelationResultModel() 从对象池中获取AlibabaAlihealthDrugKytDrSinglerelationResultModel
func GetAlibabaAlihealthDrugKytDrSinglerelationResultModel() *AlibabaAlihealthDrugKytDrSinglerelationResultModel {
	return poolAlibabaAlihealthDrugKytDrSinglerelationResultModel.Get().(*AlibabaAlihealthDrugKytDrSinglerelationResultModel)
}

// ReleaseAlibabaAlihealthDrugKytDrSinglerelationResultModel 释放AlibabaAlihealthDrugKytDrSinglerelationResultModel
func ReleaseAlibabaAlihealthDrugKytDrSinglerelationResultModel(v *AlibabaAlihealthDrugKytDrSinglerelationResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDrSinglerelationResultModel.Put(v)
}
