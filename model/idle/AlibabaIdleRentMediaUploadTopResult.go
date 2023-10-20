package idle

import (
	"sync"
)

// AlibabaIdleRentMediaUploadTopResult 结构体
type AlibabaIdleRentMediaUploadTopResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 媒体文件id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentMediaUploadTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentMediaUploadTopResult)
	},
}

// GetAlibabaIdleRentMediaUploadTopResult() 从对象池中获取AlibabaIdleRentMediaUploadTopResult
func GetAlibabaIdleRentMediaUploadTopResult() *AlibabaIdleRentMediaUploadTopResult {
	return poolAlibabaIdleRentMediaUploadTopResult.Get().(*AlibabaIdleRentMediaUploadTopResult)
}

// ReleaseAlibabaIdleRentMediaUploadTopResult 释放AlibabaIdleRentMediaUploadTopResult
func ReleaseAlibabaIdleRentMediaUploadTopResult(v *AlibabaIdleRentMediaUploadTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = 0
	v.Success = false
	poolAlibabaIdleRentMediaUploadTopResult.Put(v)
}
