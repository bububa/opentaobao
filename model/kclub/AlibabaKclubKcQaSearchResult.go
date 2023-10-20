package kclub

import (
	"sync"
)

// AlibabaKclubKcQaSearchResult 结构体
type AlibabaKclubKcQaSearchResult struct {
	// 返回数据列表
	DataList []KcSearchQuestion `json:"data_list,omitempty" xml:"data_list>kc_search_question,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaKclubKcQaSearchResult = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQaSearchResult)
	},
}

// GetAlibabaKclubKcQaSearchResult() 从对象池中获取AlibabaKclubKcQaSearchResult
func GetAlibabaKclubKcQaSearchResult() *AlibabaKclubKcQaSearchResult {
	return poolAlibabaKclubKcQaSearchResult.Get().(*AlibabaKclubKcQaSearchResult)
}

// ReleaseAlibabaKclubKcQaSearchResult 释放AlibabaKclubKcQaSearchResult
func ReleaseAlibabaKclubKcQaSearchResult(v *AlibabaKclubKcQaSearchResult) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaKclubKcQaSearchResult.Put(v)
}
