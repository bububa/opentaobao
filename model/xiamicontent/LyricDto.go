package xiamicontent

import (
	"sync"
)

// LyricDto 结构体
type LyricDto struct {
	// 歌词地址
	LyricUrl string `json:"lyric_url,omitempty" xml:"lyric_url,omitempty"`
	// 歌词类型 TXT LRC TRC TLRC TTRC
	LyricType string `json:"lyric_type,omitempty" xml:"lyric_type,omitempty"`
	// 歌词id
	LyricId int64 `json:"lyric_id,omitempty" xml:"lyric_id,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 歌曲id
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
}

var poolLyricDto = sync.Pool{
	New: func() any {
		return new(LyricDto)
	},
}

// GetLyricDto() 从对象池中获取LyricDto
func GetLyricDto() *LyricDto {
	return poolLyricDto.Get().(*LyricDto)
}

// ReleaseLyricDto 释放LyricDto
func ReleaseLyricDto(v *LyricDto) {
	v.LyricUrl = ""
	v.LyricType = ""
	v.LyricId = 0
	v.Type = 0
	v.SongId = 0
	poolLyricDto.Put(v)
}
