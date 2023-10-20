package wdk

import (
	"sync"
)

// AlibabaWdkItemStoreskuUpdateResult 结构体
type AlibabaWdkItemStoreskuUpdateResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemStoreskuUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemStoreskuUpdateResult)
	},
}

// GetAlibabaWdkItemStoreskuUpdateResult() 从对象池中获取AlibabaWdkItemStoreskuUpdateResult
func GetAlibabaWdkItemStoreskuUpdateResult() *AlibabaWdkItemStoreskuUpdateResult {
	return poolAlibabaWdkItemStoreskuUpdateResult.Get().(*AlibabaWdkItemStoreskuUpdateResult)
}

// ReleaseAlibabaWdkItemStoreskuUpdateResult 释放AlibabaWdkItemStoreskuUpdateResult
func ReleaseAlibabaWdkItemStoreskuUpdateResult(v *AlibabaWdkItemStoreskuUpdateResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemStoreskuUpdateResult.Put(v)
}
