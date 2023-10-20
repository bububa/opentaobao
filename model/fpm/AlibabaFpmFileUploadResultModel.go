package fpm

import (
	"sync"
)

// AlibabaFpmFileUploadResultModel 结构体
type AlibabaFpmFileUploadResultModel struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错信信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// returnValue
	ReturnValue *FileUploadReponseDto `json:"return_value,omitempty" xml:"return_value,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaFpmFileUploadResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaFpmFileUploadResultModel)
	},
}

// GetAlibabaFpmFileUploadResultModel() 从对象池中获取AlibabaFpmFileUploadResultModel
func GetAlibabaFpmFileUploadResultModel() *AlibabaFpmFileUploadResultModel {
	return poolAlibabaFpmFileUploadResultModel.Get().(*AlibabaFpmFileUploadResultModel)
}

// ReleaseAlibabaFpmFileUploadResultModel 释放AlibabaFpmFileUploadResultModel
func ReleaseAlibabaFpmFileUploadResultModel(v *AlibabaFpmFileUploadResultModel) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.ReturnValue = nil
	v.Success = false
	poolAlibabaFpmFileUploadResultModel.Put(v)
}
