package xiami

// TagGenreSongresult 
type TagGenreSongresult struct {
    // 总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 是否还有
    More   bool `json:"more,omitempty" xml:"more,omitempty"`
    // 歌曲列表
    Songs   []Songs `json:"songs,omitempty" xml:"songs>songs,omitempty"`
}
