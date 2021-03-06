package xiamicontent

// Page 结构体
type Page struct {
	// 歌曲列表
	SongList []SongInfoDto `json:"song_list,omitempty" xml:"song_list>song_info_dto,omitempty"`
	// 分页信息
	Paging *PagingVo `json:"paging,omitempty" xml:"paging,omitempty"`
}
