package icburfq

import (
	"sync"
)

// PageView 结构体
type PageView struct {
	// 当前页面
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 推荐数量
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
	// 总页数
	TotalPages int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
}

var poolPageView = sync.Pool{
	New: func() any {
		return new(PageView)
	},
}

// GetPageView() 从对象池中获取PageView
func GetPageView() *PageView {
	return poolPageView.Get().(*PageView)
}

// ReleasePageView 释放PageView
func ReleasePageView(v *PageView) {
	v.Current = 0
	v.PageSize = 0
	v.TotalItem = 0
	v.TotalPages = 0
	poolPageView.Put(v)
}
