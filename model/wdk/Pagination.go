package wdk

import (
	"sync"
)

// Pagination 结构体
type Pagination struct {
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 页容量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolPagination = sync.Pool{
	New: func() any {
		return new(Pagination)
	},
}

// GetPagination() 从对象池中获取Pagination
func GetPagination() *Pagination {
	return poolPagination.Get().(*Pagination)
}

// ReleasePagination 释放Pagination
func ReleasePagination(v *Pagination) {
	v.TotalPage = 0
	v.TotalCount = 0
	v.PageSize = 0
	v.CurrentPage = 0
	poolPagination.Put(v)
}
