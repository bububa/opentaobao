package kclub

import (
	"sync"
)

// AlibabaKclubKcQaGetResult 结构体
type AlibabaKclubKcQaGetResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数据
	Data *KcQaRead `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaKclubKcQaGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQaGetResult)
	},
}

// GetAlibabaKclubKcQaGetResult() 从对象池中获取AlibabaKclubKcQaGetResult
func GetAlibabaKclubKcQaGetResult() *AlibabaKclubKcQaGetResult {
	return poolAlibabaKclubKcQaGetResult.Get().(*AlibabaKclubKcQaGetResult)
}

// ReleaseAlibabaKclubKcQaGetResult 释放AlibabaKclubKcQaGetResult
func ReleaseAlibabaKclubKcQaGetResult(v *AlibabaKclubKcQaGetResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaKclubKcQaGetResult.Put(v)
}
