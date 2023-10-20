package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytSinglerelationResultModel 结构体
type AlibabaAlihealthDrugKytSinglerelationResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytSinglerelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSinglerelationResultModel)
	},
}

// GetAlibabaAlihealthDrugKytSinglerelationResultModel() 从对象池中获取AlibabaAlihealthDrugKytSinglerelationResultModel
func GetAlibabaAlihealthDrugKytSinglerelationResultModel() *AlibabaAlihealthDrugKytSinglerelationResultModel {
	return poolAlibabaAlihealthDrugKytSinglerelationResultModel.Get().(*AlibabaAlihealthDrugKytSinglerelationResultModel)
}

// ReleaseAlibabaAlihealthDrugKytSinglerelationResultModel 释放AlibabaAlihealthDrugKytSinglerelationResultModel
func ReleaseAlibabaAlihealthDrugKytSinglerelationResultModel(v *AlibabaAlihealthDrugKytSinglerelationResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytSinglerelationResultModel.Put(v)
}
