package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqyUploadrelationResultModel 结构体
type AlibabaAlihealthDrugKytScqyUploadrelationResultModel struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *SaveCodeRelationResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqyUploadrelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyUploadrelationResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqyUploadrelationResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqyUploadrelationResultModel
func GetAlibabaAlihealthDrugKytScqyUploadrelationResultModel() *AlibabaAlihealthDrugKytScqyUploadrelationResultModel {
	return poolAlibabaAlihealthDrugKytScqyUploadrelationResultModel.Get().(*AlibabaAlihealthDrugKytScqyUploadrelationResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqyUploadrelationResultModel 释放AlibabaAlihealthDrugKytScqyUploadrelationResultModel
func ReleaseAlibabaAlihealthDrugKytScqyUploadrelationResultModel(v *AlibabaAlihealthDrugKytScqyUploadrelationResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytScqyUploadrelationResultModel.Put(v)
}
