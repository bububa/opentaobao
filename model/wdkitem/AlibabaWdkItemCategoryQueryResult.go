package wdkitem

import (
	"sync"
)

// AlibabaWdkItemCategoryQueryResult 结构体
type AlibabaWdkItemCategoryQueryResult struct {
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

var poolAlibabaWdkItemCategoryQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCategoryQueryResult)
	},
}

// GetAlibabaWdkItemCategoryQueryResult() 从对象池中获取AlibabaWdkItemCategoryQueryResult
func GetAlibabaWdkItemCategoryQueryResult() *AlibabaWdkItemCategoryQueryResult {
	return poolAlibabaWdkItemCategoryQueryResult.Get().(*AlibabaWdkItemCategoryQueryResult)
}

// ReleaseAlibabaWdkItemCategoryQueryResult 释放AlibabaWdkItemCategoryQueryResult
func ReleaseAlibabaWdkItemCategoryQueryResult(v *AlibabaWdkItemCategoryQueryResult) {
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Result = ""
	v.Success = false
	poolAlibabaWdkItemCategoryQueryResult.Put(v)
}
