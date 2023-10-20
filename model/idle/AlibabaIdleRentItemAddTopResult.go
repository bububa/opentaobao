package idle

import (
	"sync"
)

// AlibabaIdleRentItemAddTopResult 结构体
type AlibabaIdleRentItemAddTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 商品id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentItemAddTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemAddTopResult)
	},
}

// GetAlibabaIdleRentItemAddTopResult() 从对象池中获取AlibabaIdleRentItemAddTopResult
func GetAlibabaIdleRentItemAddTopResult() *AlibabaIdleRentItemAddTopResult {
	return poolAlibabaIdleRentItemAddTopResult.Get().(*AlibabaIdleRentItemAddTopResult)
}

// ReleaseAlibabaIdleRentItemAddTopResult 释放AlibabaIdleRentItemAddTopResult
func ReleaseAlibabaIdleRentItemAddTopResult(v *AlibabaIdleRentItemAddTopResult) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Data = 0
	v.Success = false
	poolAlibabaIdleRentItemAddTopResult.Put(v)
}
