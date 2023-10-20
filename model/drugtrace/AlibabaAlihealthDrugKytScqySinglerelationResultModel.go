package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqySinglerelationResultModel 结构体
type AlibabaAlihealthDrugKytScqySinglerelationResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqySinglerelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqySinglerelationResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqySinglerelationResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqySinglerelationResultModel
func GetAlibabaAlihealthDrugKytScqySinglerelationResultModel() *AlibabaAlihealthDrugKytScqySinglerelationResultModel {
	return poolAlibabaAlihealthDrugKytScqySinglerelationResultModel.Get().(*AlibabaAlihealthDrugKytScqySinglerelationResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqySinglerelationResultModel 释放AlibabaAlihealthDrugKytScqySinglerelationResultModel
func ReleaseAlibabaAlihealthDrugKytScqySinglerelationResultModel(v *AlibabaAlihealthDrugKytScqySinglerelationResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytScqySinglerelationResultModel.Put(v)
}
