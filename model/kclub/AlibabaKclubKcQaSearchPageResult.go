package kclub

import (
	"sync"
)

// AlibabaKclubKcQaSearchPageResult 结构体
type AlibabaKclubKcQaSearchPageResult struct {
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页数据
	Data *Paging `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaKclubKcQaSearchPageResult = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQaSearchPageResult)
	},
}

// GetAlibabaKclubKcQaSearchPageResult() 从对象池中获取AlibabaKclubKcQaSearchPageResult
func GetAlibabaKclubKcQaSearchPageResult() *AlibabaKclubKcQaSearchPageResult {
	return poolAlibabaKclubKcQaSearchPageResult.Get().(*AlibabaKclubKcQaSearchPageResult)
}

// ReleaseAlibabaKclubKcQaSearchPageResult 释放AlibabaKclubKcQaSearchPageResult
func ReleaseAlibabaKclubKcQaSearchPageResult(v *AlibabaKclubKcQaSearchPageResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolAlibabaKclubKcQaSearchPageResult.Put(v)
}
