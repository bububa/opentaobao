package idle

import (
	"sync"
)

// AlibabaIdleRentOrderCheckstatusUploadTopResult 结构体
type AlibabaIdleRentOrderCheckstatusUploadTopResult struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentOrderCheckstatusUploadTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderCheckstatusUploadTopResult)
	},
}

// GetAlibabaIdleRentOrderCheckstatusUploadTopResult() 从对象池中获取AlibabaIdleRentOrderCheckstatusUploadTopResult
func GetAlibabaIdleRentOrderCheckstatusUploadTopResult() *AlibabaIdleRentOrderCheckstatusUploadTopResult {
	return poolAlibabaIdleRentOrderCheckstatusUploadTopResult.Get().(*AlibabaIdleRentOrderCheckstatusUploadTopResult)
}

// ReleaseAlibabaIdleRentOrderCheckstatusUploadTopResult 释放AlibabaIdleRentOrderCheckstatusUploadTopResult
func ReleaseAlibabaIdleRentOrderCheckstatusUploadTopResult(v *AlibabaIdleRentOrderCheckstatusUploadTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentOrderCheckstatusUploadTopResult.Put(v)
}
