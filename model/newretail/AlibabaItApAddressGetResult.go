package newretail

import (
	"sync"
)

// AlibabaItApAddressGetResult 结构体
type AlibabaItApAddressGetResult struct {
	// 返回的错误message
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的位置结构体
	Data *ApAddressInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 返回的错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口返回成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaItApAddressGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaItApAddressGetResult)
	},
}

// GetAlibabaItApAddressGetResult() 从对象池中获取AlibabaItApAddressGetResult
func GetAlibabaItApAddressGetResult() *AlibabaItApAddressGetResult {
	return poolAlibabaItApAddressGetResult.Get().(*AlibabaItApAddressGetResult)
}

// ReleaseAlibabaItApAddressGetResult 释放AlibabaItApAddressGetResult
func ReleaseAlibabaItApAddressGetResult(v *AlibabaItApAddressGetResult) {
	v.ErrorMsg = ""
	v.Data = nil
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaItApAddressGetResult.Put(v)
}
