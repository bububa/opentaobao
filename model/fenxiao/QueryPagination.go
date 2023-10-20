package fenxiao

import (
	"sync"
)

// QueryPagination 结构体
type QueryPagination struct {
	// 当前页码数
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolQueryPagination = sync.Pool{
	New: func() any {
		return new(QueryPagination)
	},
}

// GetQueryPagination() 从对象池中获取QueryPagination
func GetQueryPagination() *QueryPagination {
	return poolQueryPagination.Get().(*QueryPagination)
}

// ReleaseQueryPagination 释放QueryPagination
func ReleaseQueryPagination(v *QueryPagination) {
	v.PageIndex = 0
	v.PageSize = 0
	poolQueryPagination.Put(v)
}
