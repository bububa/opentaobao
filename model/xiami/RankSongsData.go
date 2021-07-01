package xiami

// RankSongsData 结构体
type RankSongsData struct {
	// 歌曲列表
	Songs []RankSong `json:"songs,omitempty" xml:"songs>rank_song,omitempty"`
}
