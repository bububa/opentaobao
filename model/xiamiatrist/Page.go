package xiamiatrist

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 艺人信息列表
	Artists []ArtistDto `json:"artists,omitempty" xml:"artists>artist_dto,omitempty"`
	// 满足条件的总数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 分页信息
	Paging *Paging `json:"paging,omitempty" xml:"paging,omitempty"`
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
	v.Artists = v.Artists[:0]
	v.Count = 0
	v.Paging = nil
	poolPage.Put(v)
}
