package newretail

import (
	"sync"
)

// AlibabaItApAddressSetResult 结构体
type AlibabaItApAddressSetResult struct {
	// 返回内容
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回的message
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaItApAddressSetResult = sync.Pool{
	New: func() any {
		return new(AlibabaItApAddressSetResult)
	},
}

// GetAlibabaItApAddressSetResult() 从对象池中获取AlibabaItApAddressSetResult
func GetAlibabaItApAddressSetResult() *AlibabaItApAddressSetResult {
	return poolAlibabaItApAddressSetResult.Get().(*AlibabaItApAddressSetResult)
}

// ReleaseAlibabaItApAddressSetResult 释放AlibabaItApAddressSetResult
func ReleaseAlibabaItApAddressSetResult(v *AlibabaItApAddressSetResult) {
	v.Data = ""
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaItApAddressSetResult.Put(v)
}
