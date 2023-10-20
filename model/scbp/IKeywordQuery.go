package scbp

import (
	"sync"
)

// IKeywordQuery 结构体
type IKeywordQuery struct {
	// 结束时间 当inteval=7或者30的时候不需要填写
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 开始时间 当inteval=7或者30的时候不需要填写
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 区间 只能为1 7 30
	Inteval int64 `json:"inteval,omitempty" xml:"inteval,omitempty"`
	// 每页行数
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
}

var poolIKeywordQuery = sync.Pool{
	New: func() any {
		return new(IKeywordQuery)
	},
}

// GetIKeywordQuery() 从对象池中获取IKeywordQuery
func GetIKeywordQuery() *IKeywordQuery {
	return poolIKeywordQuery.Get().(*IKeywordQuery)
}

// ReleaseIKeywordQuery 释放IKeywordQuery
func ReleaseIKeywordQuery(v *IKeywordQuery) {
	v.EndDate = ""
	v.Keyword = ""
	v.BeginDate = ""
	v.Inteval = 0
	v.PerPageSize = 0
	v.ToPage = 0
	poolIKeywordQuery.Put(v)
}
