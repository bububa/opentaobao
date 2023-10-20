package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMerchantstoreskuUpdateResult 结构体
type AlibabaWdkItemMerchantstoreskuUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMerchantstoreskuUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantstoreskuUpdateResult)
	},
}

// GetAlibabaWdkItemMerchantstoreskuUpdateResult() 从对象池中获取AlibabaWdkItemMerchantstoreskuUpdateResult
func GetAlibabaWdkItemMerchantstoreskuUpdateResult() *AlibabaWdkItemMerchantstoreskuUpdateResult {
	return poolAlibabaWdkItemMerchantstoreskuUpdateResult.Get().(*AlibabaWdkItemMerchantstoreskuUpdateResult)
}

// ReleaseAlibabaWdkItemMerchantstoreskuUpdateResult 释放AlibabaWdkItemMerchantstoreskuUpdateResult
func ReleaseAlibabaWdkItemMerchantstoreskuUpdateResult(v *AlibabaWdkItemMerchantstoreskuUpdateResult) {
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemMerchantstoreskuUpdateResult.Put(v)
}
