package drugtrace

import (
	"sync"
)

// AlibabaAlihealthZyUploadrelationResultModel 结构体
type AlibabaAlihealthZyUploadrelationResultModel struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *SaveCodeRelationResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthZyUploadrelationResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthZyUploadrelationResultModel)
	},
}

// GetAlibabaAlihealthZyUploadrelationResultModel() 从对象池中获取AlibabaAlihealthZyUploadrelationResultModel
func GetAlibabaAlihealthZyUploadrelationResultModel() *AlibabaAlihealthZyUploadrelationResultModel {
	return poolAlibabaAlihealthZyUploadrelationResultModel.Get().(*AlibabaAlihealthZyUploadrelationResultModel)
}

// ReleaseAlibabaAlihealthZyUploadrelationResultModel 释放AlibabaAlihealthZyUploadrelationResultModel
func ReleaseAlibabaAlihealthZyUploadrelationResultModel(v *AlibabaAlihealthZyUploadrelationResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.ResponseSuccess = false
	poolAlibabaAlihealthZyUploadrelationResultModel.Put(v)
}
