package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel 结构体
type AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel struct {
	// 内层大对象
	Models []CodeDrugInfoDto `json:"models,omitempty" xml:"models>code_drug_info_dto,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口调用标识：true/false
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel() 从对象池中获取AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel
func GetAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel() *AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel {
	return poolAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel.Get().(*AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel 释放AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel
func ReleaseAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel(v *AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel) {
	v.Models = v.Models[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugCodeKytYqQuerycodeResultModel.Put(v)
}
