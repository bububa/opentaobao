package scbp

import (
	"sync"
)

// RecKeywordQuery 结构体
type RecKeywordQuery struct {
	// 搜索词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 每页行数
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
}

var poolRecKeywordQuery = sync.Pool{
	New: func() any {
		return new(RecKeywordQuery)
	},
}

// GetRecKeywordQuery() 从对象池中获取RecKeywordQuery
func GetRecKeywordQuery() *RecKeywordQuery {
	return poolRecKeywordQuery.Get().(*RecKeywordQuery)
}

// ReleaseRecKeywordQuery 释放RecKeywordQuery
func ReleaseRecKeywordQuery(v *RecKeywordQuery) {
	v.Keyword = ""
	v.PerPageSize = 0
	v.ToPage = 0
	poolRecKeywordQuery.Put(v)
}
