package dmp

import (
	"sync"
)

// Pager 结构体
type Pager struct {
	// 当前页数
	IntCurrentPage int64 `json:"int_current_page,omitempty" xml:"int_current_page,omitempty"`
	// 分页大小
	IntPageSize int64 `json:"int_page_size,omitempty" xml:"int_page_size,omitempty"`
	// 记录总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolPager = sync.Pool{
	New: func() any {
		return new(Pager)
	},
}

// GetPager() 从对象池中获取Pager
func GetPager() *Pager {
	return poolPager.Get().(*Pager)
}

// ReleasePager 释放Pager
func ReleasePager(v *Pager) {
	v.IntCurrentPage = 0
	v.IntPageSize = 0
	v.Total = 0
	v.PageSize = 0
	v.CurrentPage = 0
	poolPager.Put(v)
}
