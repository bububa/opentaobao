package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMorebarcodeOpsResult 结构体
type AlibabaWdkItemMorebarcodeOpsResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMorebarcodeOpsResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMorebarcodeOpsResult)
	},
}

// GetAlibabaWdkItemMorebarcodeOpsResult() 从对象池中获取AlibabaWdkItemMorebarcodeOpsResult
func GetAlibabaWdkItemMorebarcodeOpsResult() *AlibabaWdkItemMorebarcodeOpsResult {
	return poolAlibabaWdkItemMorebarcodeOpsResult.Get().(*AlibabaWdkItemMorebarcodeOpsResult)
}

// ReleaseAlibabaWdkItemMorebarcodeOpsResult 释放AlibabaWdkItemMorebarcodeOpsResult
func ReleaseAlibabaWdkItemMorebarcodeOpsResult(v *AlibabaWdkItemMorebarcodeOpsResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemMorebarcodeOpsResult.Put(v)
}
