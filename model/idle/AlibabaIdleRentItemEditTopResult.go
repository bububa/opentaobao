package idle

import (
	"sync"
)

// AlibabaIdleRentItemEditTopResult 结构体
type AlibabaIdleRentItemEditTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 是否更新成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentItemEditTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemEditTopResult)
	},
}

// GetAlibabaIdleRentItemEditTopResult() 从对象池中获取AlibabaIdleRentItemEditTopResult
func GetAlibabaIdleRentItemEditTopResult() *AlibabaIdleRentItemEditTopResult {
	return poolAlibabaIdleRentItemEditTopResult.Get().(*AlibabaIdleRentItemEditTopResult)
}

// ReleaseAlibabaIdleRentItemEditTopResult 释放AlibabaIdleRentItemEditTopResult
func ReleaseAlibabaIdleRentItemEditTopResult(v *AlibabaIdleRentItemEditTopResult) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentItemEditTopResult.Put(v)
}
