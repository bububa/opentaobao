package alihealthoutflow

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 页对象
	Pages []DrugDto `json:"pages,omitempty" xml:"pages>drug_dto,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
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
	v.Pages = v.Pages[:0]
	v.TotalCount = 0
	v.PageSize = 0
	v.PageNo = 0
	poolPage.Put(v)
}
