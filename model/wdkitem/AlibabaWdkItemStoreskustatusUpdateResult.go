package wdkitem

import (
	"sync"
)

// AlibabaWdkItemStoreskustatusUpdateResult 结构体
type AlibabaWdkItemStoreskustatusUpdateResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemStoreskustatusUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemStoreskustatusUpdateResult)
	},
}

// GetAlibabaWdkItemStoreskustatusUpdateResult() 从对象池中获取AlibabaWdkItemStoreskustatusUpdateResult
func GetAlibabaWdkItemStoreskustatusUpdateResult() *AlibabaWdkItemStoreskustatusUpdateResult {
	return poolAlibabaWdkItemStoreskustatusUpdateResult.Get().(*AlibabaWdkItemStoreskustatusUpdateResult)
}

// ReleaseAlibabaWdkItemStoreskustatusUpdateResult 释放AlibabaWdkItemStoreskustatusUpdateResult
func ReleaseAlibabaWdkItemStoreskustatusUpdateResult(v *AlibabaWdkItemStoreskustatusUpdateResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemStoreskustatusUpdateResult.Put(v)
}
