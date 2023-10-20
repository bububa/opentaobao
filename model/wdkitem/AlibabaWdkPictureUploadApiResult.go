package wdkitem

import (
	"sync"
)

// AlibabaWdkPictureUploadApiResult 结构体
type AlibabaWdkPictureUploadApiResult struct {
	// 错误code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// model
	Model *PictureDo `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkPictureUploadApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPictureUploadApiResult)
	},
}

// GetAlibabaWdkPictureUploadApiResult() 从对象池中获取AlibabaWdkPictureUploadApiResult
func GetAlibabaWdkPictureUploadApiResult() *AlibabaWdkPictureUploadApiResult {
	return poolAlibabaWdkPictureUploadApiResult.Get().(*AlibabaWdkPictureUploadApiResult)
}

// ReleaseAlibabaWdkPictureUploadApiResult 释放AlibabaWdkPictureUploadApiResult
func ReleaseAlibabaWdkPictureUploadApiResult(v *AlibabaWdkPictureUploadApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkPictureUploadApiResult.Put(v)
}
