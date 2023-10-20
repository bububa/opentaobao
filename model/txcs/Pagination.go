package txcs

import (
	"sync"
)

// Pagination 结构体
type Pagination struct {
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 单页个数
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
	v.Total = 0
	v.PageSize = 0
	v.CurrentPage = 0
	poolPagination.Put(v)
}
