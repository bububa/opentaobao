package shenjing

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 活动列表
	Items []AlibabaShenjingCoreActivityGetappshowlistT `json:"items,omitempty" xml:"items>alibaba_shenjing_core_activity_getappshowlist_t,omitempty"`
	// 分页总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 一页行数
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 分页总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 一页行数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

var poolPage = sync.Pool{
	New: func() any {
		return new(Page)
	},
}

// GetPage() 从对象池中获取Page
func GetPage() *Page {
	return poolPage.Get().(*Page)
}

// ReleasePage 释放Page
func ReleasePage(v *Page) {
	v.Items = v.Items[:0]
	v.Total = 0
	v.Limit = 0
	v.TotalCount = 0
	v.PageSize = 0
	v.CurrentPage = 0
	v.TotalPage = 0
	poolPage.Put(v)
}
