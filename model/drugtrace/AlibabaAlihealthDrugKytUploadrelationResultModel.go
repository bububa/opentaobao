package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytUploadrelationResultModel 结构体
type AlibabaAlihealthDrugKytUploadrelationResultModel struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *SaveCodeRelationResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytUploadrelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUploadrelationResultModel)
	},
}

// GetAlibabaAlihealthDrugKytUploadrelationResultModel() 从对象池中获取AlibabaAlihealthDrugKytUploadrelationResultModel
func GetAlibabaAlihealthDrugKytUploadrelationResultModel() *AlibabaAlihealthDrugKytUploadrelationResultModel {
	return poolAlibabaAlihealthDrugKytUploadrelationResultModel.Get().(*AlibabaAlihealthDrugKytUploadrelationResultModel)
}

// ReleaseAlibabaAlihealthDrugKytUploadrelationResultModel 释放AlibabaAlihealthDrugKytUploadrelationResultModel
func ReleaseAlibabaAlihealthDrugKytUploadrelationResultModel(v *AlibabaAlihealthDrugKytUploadrelationResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytUploadrelationResultModel.Put(v)
}
