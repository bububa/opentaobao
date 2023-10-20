package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel
func GetAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel() *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel 释放AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel(v *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel.Put(v)
}
