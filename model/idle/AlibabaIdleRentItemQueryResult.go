package idle

import (
	"sync"
)

// AlibabaIdleRentItemQueryResult 结构体
type AlibabaIdleRentItemQueryResult struct {
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回素材id
	Data *AlibabaIdleRentItemQueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 数据是否可用
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentItemQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemQueryResult)
	},
}

// GetAlibabaIdleRentItemQueryResult() 从对象池中获取AlibabaIdleRentItemQueryResult
func GetAlibabaIdleRentItemQueryResult() *AlibabaIdleRentItemQueryResult {
	return poolAlibabaIdleRentItemQueryResult.Get().(*AlibabaIdleRentItemQueryResult)
}

// ReleaseAlibabaIdleRentItemQueryResult 释放AlibabaIdleRentItemQueryResult
func ReleaseAlibabaIdleRentItemQueryResult(v *AlibabaIdleRentItemQueryResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaIdleRentItemQueryResult.Put(v)
}
