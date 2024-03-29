package wdkitem

import (
	"sync"
)

// AlibabaWdkItemBrandQueryResult 结构体
type AlibabaWdkItemBrandQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemBrandQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemBrandQueryResult)
	},
}

// GetAlibabaWdkItemBrandQueryResult() 从对象池中获取AlibabaWdkItemBrandQueryResult
func GetAlibabaWdkItemBrandQueryResult() *AlibabaWdkItemBrandQueryResult {
	return poolAlibabaWdkItemBrandQueryResult.Get().(*AlibabaWdkItemBrandQueryResult)
}

// ReleaseAlibabaWdkItemBrandQueryResult 释放AlibabaWdkItemBrandQueryResult
func ReleaseAlibabaWdkItemBrandQueryResult(v *AlibabaWdkItemBrandQueryResult) {
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Result = ""
	v.Success = false
	poolAlibabaWdkItemBrandQueryResult.Put(v)
}
