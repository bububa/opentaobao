package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMerchantstoreskuCreateResult 结构体
type AlibabaWdkItemMerchantstoreskuCreateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMerchantstoreskuCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantstoreskuCreateResult)
	},
}

// GetAlibabaWdkItemMerchantstoreskuCreateResult() 从对象池中获取AlibabaWdkItemMerchantstoreskuCreateResult
func GetAlibabaWdkItemMerchantstoreskuCreateResult() *AlibabaWdkItemMerchantstoreskuCreateResult {
	return poolAlibabaWdkItemMerchantstoreskuCreateResult.Get().(*AlibabaWdkItemMerchantstoreskuCreateResult)
}

// ReleaseAlibabaWdkItemMerchantstoreskuCreateResult 释放AlibabaWdkItemMerchantstoreskuCreateResult
func ReleaseAlibabaWdkItemMerchantstoreskuCreateResult(v *AlibabaWdkItemMerchantstoreskuCreateResult) {
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemMerchantstoreskuCreateResult.Put(v)
}
