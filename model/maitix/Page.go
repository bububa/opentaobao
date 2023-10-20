package maitix

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 数据对象
	DataArrList []ProjectDto `json:"data_arr_list,omitempty" xml:"data_arr_list>project_dto,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总项目数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
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
	v.DataArrList = v.DataArrList[:0]
	v.CurrentPage = 0
	v.PageCount = 0
	v.PageSize = 0
	v.TotalCount = 0
	poolPage.Put(v)
}
