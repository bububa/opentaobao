package xiamicontent

// Page 结构体
type Page struct {
	// 音乐列表
	Data []MusicDto `json:"data,omitempty" xml:"data>music_dto,omitempty"`
	// 歌曲列表
	SongList []SongInfoDto `json:"song_list,omitempty" xml:"song_list>song_info_dto,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 分页信息
	PagingVo *PagingVO `json:"paging_vo,omitempty" xml:"paging_vo,omitempty"`
	// 分页信息
	Paging *PagingVo `json:"paging,omitempty" xml:"paging,omitempty"`
}
