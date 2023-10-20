package lsttrade

import (
	"sync"
)

// LstOrderQuery 结构体
type LstOrderQuery struct {
	// 开始时间，支持查询最近一年到最近一小时，[begin,end)，最大半小时
	GmtCreateBegin string `json:"gmt_create_begin,omitempty" xml:"gmt_create_begin,omitempty"`
	// 结束时间，参考开始时间
	GmtCreateEnd string `json:"gmt_create_end,omitempty" xml:"gmt_create_end,omitempty"`
	// 页码，从1开始
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolLstOrderQuery = sync.Pool{
	New: func() any {
		return new(LstOrderQuery)
	},
}

// GetLstOrderQuery() 从对象池中获取LstOrderQuery
func GetLstOrderQuery() *LstOrderQuery {
	return poolLstOrderQuery.Get().(*LstOrderQuery)
}

// ReleaseLstOrderQuery 释放LstOrderQuery
func ReleaseLstOrderQuery(v *LstOrderQuery) {
	v.GmtCreateBegin = ""
	v.GmtCreateEnd = ""
	v.Page = 0
	v.Size = 0
	poolLstOrderQuery.Put(v)
}
