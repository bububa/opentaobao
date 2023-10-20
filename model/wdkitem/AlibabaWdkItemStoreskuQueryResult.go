package wdkitem

import (
	"sync"
)

// AlibabaWdkItemStoreskuQueryResult 结构体
type AlibabaWdkItemStoreskuQueryResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemStoreskuQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemStoreskuQueryResult)
	},
}

// GetAlibabaWdkItemStoreskuQueryResult() 从对象池中获取AlibabaWdkItemStoreskuQueryResult
func GetAlibabaWdkItemStoreskuQueryResult() *AlibabaWdkItemStoreskuQueryResult {
	return poolAlibabaWdkItemStoreskuQueryResult.Get().(*AlibabaWdkItemStoreskuQueryResult)
}

// ReleaseAlibabaWdkItemStoreskuQueryResult 释放AlibabaWdkItemStoreskuQueryResult
func ReleaseAlibabaWdkItemStoreskuQueryResult(v *AlibabaWdkItemStoreskuQueryResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Code = ""
	v.Result = ""
	v.Success = false
	poolAlibabaWdkItemStoreskuQueryResult.Put(v)
}
