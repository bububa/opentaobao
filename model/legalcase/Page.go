package legalcase

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 返回列表
	Datas []StandpointSearchDto `json:"datas,omitempty" xml:"datas>standpoint_search_dto,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总数
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
	v.Datas = v.Datas[:0]
	v.CurrentPage = 0
	v.TotalCount = 0
	poolPage.Put(v)
}
