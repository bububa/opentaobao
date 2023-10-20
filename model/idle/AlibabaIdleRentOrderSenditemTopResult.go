package idle

import (
	"sync"
)

// AlibabaIdleRentOrderSenditemTopResult 结构体
type AlibabaIdleRentOrderSenditemTopResult struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentOrderSenditemTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderSenditemTopResult)
	},
}

// GetAlibabaIdleRentOrderSenditemTopResult() 从对象池中获取AlibabaIdleRentOrderSenditemTopResult
func GetAlibabaIdleRentOrderSenditemTopResult() *AlibabaIdleRentOrderSenditemTopResult {
	return poolAlibabaIdleRentOrderSenditemTopResult.Get().(*AlibabaIdleRentOrderSenditemTopResult)
}

// ReleaseAlibabaIdleRentOrderSenditemTopResult 释放AlibabaIdleRentOrderSenditemTopResult
func ReleaseAlibabaIdleRentOrderSenditemTopResult(v *AlibabaIdleRentOrderSenditemTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentOrderSenditemTopResult.Put(v)
}
