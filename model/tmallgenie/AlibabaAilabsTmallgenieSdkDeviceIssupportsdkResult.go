package tmallgenie

import (
	"sync"
)

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult 结构体
type AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否支持
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult)
	},
}

// GetAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult() 从对象池中获取AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult
func GetAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult() *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult {
	return poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult.Get().(*AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult)
}

// ReleaseAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult 释放AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult
func ReleaseAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult(v *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult) {
	v.Message = ""
	v.Code = 0
	v.Result = false
	poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult.Put(v)
}
