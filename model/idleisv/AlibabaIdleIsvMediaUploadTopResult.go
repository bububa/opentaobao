package idleisv

import (
	"sync"
)

// AlibabaIdleIsvMediaUploadTopResult 结构体
type AlibabaIdleIsvMediaUploadTopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 媒体文件id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvMediaUploadTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvMediaUploadTopResult)
	},
}

// GetAlibabaIdleIsvMediaUploadTopResult() 从对象池中获取AlibabaIdleIsvMediaUploadTopResult
func GetAlibabaIdleIsvMediaUploadTopResult() *AlibabaIdleIsvMediaUploadTopResult {
	return poolAlibabaIdleIsvMediaUploadTopResult.Get().(*AlibabaIdleIsvMediaUploadTopResult)
}

// ReleaseAlibabaIdleIsvMediaUploadTopResult 释放AlibabaIdleIsvMediaUploadTopResult
func ReleaseAlibabaIdleIsvMediaUploadTopResult(v *AlibabaIdleIsvMediaUploadTopResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = 0
	v.Success = false
	poolAlibabaIdleIsvMediaUploadTopResult.Put(v)
}
