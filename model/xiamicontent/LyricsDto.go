package xiamicontent

// LyricsDTO 
type LyricsDTO struct {
    // 歌词文件地址
    LyricUrl   string `json:"lyric_url,omitempty" xml:"lyric_url,omitempty"`
    // 歌词类型:TXT(文本歌词) | LRC(逐行动态歌词) | TRC(逐字动态歌词) | TLRC(翻译LRC歌词) | OSS_TRANSLATE(翻译TRC歌词)
    LyricType   string `json:"lyric_type,omitempty" xml:"lyric_type,omitempty"`
    // 所属歌曲ID
    SongId   int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
    // 歌词ID
    LyricId   int64 `json:"lyric_id,omitempty" xml:"lyric_id,omitempty"`
}
