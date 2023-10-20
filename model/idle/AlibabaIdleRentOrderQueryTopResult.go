package idle

import (
	"sync"
)

// AlibabaIdleRentOrderQueryTopResult 结构体
type AlibabaIdleRentOrderQueryTopResult struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data *RentalOrderDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentOrderQueryTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderQueryTopResult)
	},
}

// GetAlibabaIdleRentOrderQueryTopResult() 从对象池中获取AlibabaIdleRentOrderQueryTopResult
func GetAlibabaIdleRentOrderQueryTopResult() *AlibabaIdleRentOrderQueryTopResult {
	return poolAlibabaIdleRentOrderQueryTopResult.Get().(*AlibabaIdleRentOrderQueryTopResult)
}

// ReleaseAlibabaIdleRentOrderQueryTopResult 释放AlibabaIdleRentOrderQueryTopResult
func ReleaseAlibabaIdleRentOrderQueryTopResult(v *AlibabaIdleRentOrderQueryTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaIdleRentOrderQueryTopResult.Put(v)
}
