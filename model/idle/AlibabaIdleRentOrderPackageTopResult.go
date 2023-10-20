package idle

import (
	"sync"
)

// AlibabaIdleRentOrderPackageTopResult 结构体
type AlibabaIdleRentOrderPackageTopResult struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentOrderPackageTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderPackageTopResult)
	},
}

// GetAlibabaIdleRentOrderPackageTopResult() 从对象池中获取AlibabaIdleRentOrderPackageTopResult
func GetAlibabaIdleRentOrderPackageTopResult() *AlibabaIdleRentOrderPackageTopResult {
	return poolAlibabaIdleRentOrderPackageTopResult.Get().(*AlibabaIdleRentOrderPackageTopResult)
}

// ReleaseAlibabaIdleRentOrderPackageTopResult 释放AlibabaIdleRentOrderPackageTopResult
func ReleaseAlibabaIdleRentOrderPackageTopResult(v *AlibabaIdleRentOrderPackageTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentOrderPackageTopResult.Put(v)
}
