package xiami

// HotSongsResult 结构体
type HotSongsResult struct {
	// 歌曲列表
	Songs []StandardSong `json:"songs,omitempty" xml:"songs>standard_song,omitempty"`
}
