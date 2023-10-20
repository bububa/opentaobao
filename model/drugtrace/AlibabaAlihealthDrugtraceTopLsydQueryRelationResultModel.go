package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel.Put(v)
}
