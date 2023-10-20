package legalsuit

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 未采纳口径对象
	Data []StandpointDraftOutPutDto `json:"data,omitempty" xml:"data>standpoint_draft_out_put_dto,omitempty"`
	// 1
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 10
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
	v.Data = v.Data[:0]
	v.CurrentPage = 0
	v.TotalCount = 0
	poolPage.Put(v)
}
