package idle

import (
	"sync"
)

// AlibabaIdleRentOrderReceiveitemTopResult 结构体
type AlibabaIdleRentOrderReceiveitemTopResult struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentOrderReceiveitemTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderReceiveitemTopResult)
	},
}

// GetAlibabaIdleRentOrderReceiveitemTopResult() 从对象池中获取AlibabaIdleRentOrderReceiveitemTopResult
func GetAlibabaIdleRentOrderReceiveitemTopResult() *AlibabaIdleRentOrderReceiveitemTopResult {
	return poolAlibabaIdleRentOrderReceiveitemTopResult.Get().(*AlibabaIdleRentOrderReceiveitemTopResult)
}

// ReleaseAlibabaIdleRentOrderReceiveitemTopResult 释放AlibabaIdleRentOrderReceiveitemTopResult
func ReleaseAlibabaIdleRentOrderReceiveitemTopResult(v *AlibabaIdleRentOrderReceiveitemTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentOrderReceiveitemTopResult.Put(v)
}
