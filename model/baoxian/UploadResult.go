package baoxian

import (
	"sync"
)

// UploadResult 结构体
type UploadResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// model
	Model *InsAttachmentUploadVo `json:"model,omitempty" xml:"model,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolUploadResult = sync.Pool{
	New: func() any {
		return new(UploadResult)
	},
}

// GetUploadResult() 从对象池中获取UploadResult
func GetUploadResult() *UploadResult {
	return poolUploadResult.Get().(*UploadResult)
}

// ReleaseUploadResult 释放UploadResult
func ReleaseUploadResult(v *UploadResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Model = nil
	v.IsSuccess = false
	poolUploadResult.Put(v)
}
