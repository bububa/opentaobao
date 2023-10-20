package wdk

import (
	"sync"
)

// AlibabaWdkItemServiceitemQueryResult 结构体
type AlibabaWdkItemServiceitemQueryResult struct {
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemServiceitemQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemServiceitemQueryResult)
	},
}

// GetAlibabaWdkItemServiceitemQueryResult() 从对象池中获取AlibabaWdkItemServiceitemQueryResult
func GetAlibabaWdkItemServiceitemQueryResult() *AlibabaWdkItemServiceitemQueryResult {
	return poolAlibabaWdkItemServiceitemQueryResult.Get().(*AlibabaWdkItemServiceitemQueryResult)
}

// ReleaseAlibabaWdkItemServiceitemQueryResult 释放AlibabaWdkItemServiceitemQueryResult
func ReleaseAlibabaWdkItemServiceitemQueryResult(v *AlibabaWdkItemServiceitemQueryResult) {
	v.Result = ""
	v.Success = false
	poolAlibabaWdkItemServiceitemQueryResult.Put(v)
}
