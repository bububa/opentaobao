package idle

import (
	"sync"
)

// AlibabaIdleRentItemSkuUpdateTopResult 结构体
type AlibabaIdleRentItemSkuUpdateTopResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentItemSkuUpdateTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemSkuUpdateTopResult)
	},
}

// GetAlibabaIdleRentItemSkuUpdateTopResult() 从对象池中获取AlibabaIdleRentItemSkuUpdateTopResult
func GetAlibabaIdleRentItemSkuUpdateTopResult() *AlibabaIdleRentItemSkuUpdateTopResult {
	return poolAlibabaIdleRentItemSkuUpdateTopResult.Get().(*AlibabaIdleRentItemSkuUpdateTopResult)
}

// ReleaseAlibabaIdleRentItemSkuUpdateTopResult 释放AlibabaIdleRentItemSkuUpdateTopResult
func ReleaseAlibabaIdleRentItemSkuUpdateTopResult(v *AlibabaIdleRentItemSkuUpdateTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentItemSkuUpdateTopResult.Put(v)
}
