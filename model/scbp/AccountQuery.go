package scbp

import (
	"sync"
)

// AccountQuery 结构体
type AccountQuery struct {
	// 开始时间 inteval=7或30不用指定
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 结束时间 inteval=7或30不用指定
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 区间 1代表指定时间查询 7代表最近7天 30代表最近30天
	Inteval int64 `json:"inteval,omitempty" xml:"inteval,omitempty"`
	// 每页行数
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
}

var poolAccountQuery = sync.Pool{
	New: func() any {
		return new(AccountQuery)
	},
}

// GetAccountQuery() 从对象池中获取AccountQuery
func GetAccountQuery() *AccountQuery {
	return poolAccountQuery.Get().(*AccountQuery)
}

// ReleaseAccountQuery 释放AccountQuery
func ReleaseAccountQuery(v *AccountQuery) {
	v.BeginDate = ""
	v.EndDate = ""
	v.Inteval = 0
	v.PerPageSize = 0
	v.ToPage = 0
	poolAccountQuery.Put(v)
}
