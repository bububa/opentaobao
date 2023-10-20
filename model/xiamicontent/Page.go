package xiamicontent

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 音乐列表
	Data []MusicDto `json:"data,omitempty" xml:"data>music_dto,omitempty"`
	// 歌曲列表
	SongList []SongInfoDto `json:"song_list,omitempty" xml:"song_list>song_info_dto,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 分页信息
	PagingVo *PagingVo `json:"paging_vo,omitempty" xml:"paging_vo,omitempty"`
	// 分页信息
	Paging *PagingVo `json:"paging,omitempty" xml:"paging,omitempty"`
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
	v.SongList = v.SongList[:0]
	v.Count = 0
	v.PagingVo = nil
	v.Paging = nil
	poolPage.Put(v)
}
