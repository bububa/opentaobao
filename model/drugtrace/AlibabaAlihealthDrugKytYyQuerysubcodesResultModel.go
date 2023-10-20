package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytYyQuerysubcodesResultModel 结构体
type AlibabaAlihealthDrugKytYyQuerysubcodesResultModel struct {
	// model
	ModelList []CodeRelationDto `json:"model_list,omitempty" xml:"model_list>code_relation_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytYyQuerysubcodesResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYyQuerysubcodesResultModel)
	},
}

// GetAlibabaAlihealthDrugKytYyQuerysubcodesResultModel() 从对象池中获取AlibabaAlihealthDrugKytYyQuerysubcodesResultModel
func GetAlibabaAlihealthDrugKytYyQuerysubcodesResultModel() *AlibabaAlihealthDrugKytYyQuerysubcodesResultModel {
	return poolAlibabaAlihealthDrugKytYyQuerysubcodesResultModel.Get().(*AlibabaAlihealthDrugKytYyQuerysubcodesResultModel)
}

// ReleaseAlibabaAlihealthDrugKytYyQuerysubcodesResultModel 释放AlibabaAlihealthDrugKytYyQuerysubcodesResultModel
func ReleaseAlibabaAlihealthDrugKytYyQuerysubcodesResultModel(v *AlibabaAlihealthDrugKytYyQuerysubcodesResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytYyQuerysubcodesResultModel.Put(v)
}
