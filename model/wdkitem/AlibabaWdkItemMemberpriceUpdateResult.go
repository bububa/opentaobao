package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMemberpriceUpdateResult 结构体
type AlibabaWdkItemMemberpriceUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMemberpriceUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMemberpriceUpdateResult)
	},
}

// GetAlibabaWdkItemMemberpriceUpdateResult() 从对象池中获取AlibabaWdkItemMemberpriceUpdateResult
func GetAlibabaWdkItemMemberpriceUpdateResult() *AlibabaWdkItemMemberpriceUpdateResult {
	return poolAlibabaWdkItemMemberpriceUpdateResult.Get().(*AlibabaWdkItemMemberpriceUpdateResult)
}

// ReleaseAlibabaWdkItemMemberpriceUpdateResult 释放AlibabaWdkItemMemberpriceUpdateResult
func ReleaseAlibabaWdkItemMemberpriceUpdateResult(v *AlibabaWdkItemMemberpriceUpdateResult) {
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Success = false
	poolAlibabaWdkItemMemberpriceUpdateResult.Put(v)
}
