package wdk

import (
	"sync"
)

// AlibabaWdkItemCategoryUpdateResult 结构体
type AlibabaWdkItemCategoryUpdateResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemCategoryUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCategoryUpdateResult)
	},
}

// GetAlibabaWdkItemCategoryUpdateResult() 从对象池中获取AlibabaWdkItemCategoryUpdateResult
func GetAlibabaWdkItemCategoryUpdateResult() *AlibabaWdkItemCategoryUpdateResult {
	return poolAlibabaWdkItemCategoryUpdateResult.Get().(*AlibabaWdkItemCategoryUpdateResult)
}

// ReleaseAlibabaWdkItemCategoryUpdateResult 释放AlibabaWdkItemCategoryUpdateResult
func ReleaseAlibabaWdkItemCategoryUpdateResult(v *AlibabaWdkItemCategoryUpdateResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemCategoryUpdateResult.Put(v)
}
