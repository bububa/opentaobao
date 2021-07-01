package xiamicontent

// LyricDto 
type LyricDto struct {
    // 歌词类型 TXT LRC TRC TLRC TTRC
    LyricType   string `json:"lyric_type,omitempty" xml:"lyric_type,omitempty"`
    // 歌词id
    LyricId   int64 `json:"lyric_id,omitempty" xml:"lyric_id,omitempty"`
    // 歌词地址
    LyricUrl   string `json:"lyric_url,omitempty" xml:"lyric_url,omitempty"`
}
