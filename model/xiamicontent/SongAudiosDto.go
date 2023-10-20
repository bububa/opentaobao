package xiamicontent

import (
	"sync"
)

// SongAudiosDto 结构体
type SongAudiosDto struct {
	// 音频信息
	Audios []AudioDto `json:"audios,omitempty" xml:"audios>audio_dto,omitempty"`
	// 歌曲ID
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
	// 副歌片段结束时间（单位:毫秒）
	RefrainEndTime int64 `json:"refrain_end_time,omitempty" xml:"refrain_end_time,omitempty"`
	// 副歌片段开始时间（单位:毫秒）
	RefrainStartTime int64 `json:"refrain_start_time,omitempty" xml:"refrain_start_time,omitempty"`
	// 热门片段开始时间（单位:毫秒）
	VideoPartStartTime int64 `json:"video_part_start_time,omitempty" xml:"video_part_start_time,omitempty"`
	// 热门片段结束时间（单位:毫秒）
	VideoPartEndTime int64 `json:"video_part_end_time,omitempty" xml:"video_part_end_time,omitempty"`
}

var poolSongAudiosDto = sync.Pool{
	New: func() any {
		return new(SongAudiosDto)
	},
}

// GetSongAudiosDto() 从对象池中获取SongAudiosDto
func GetSongAudiosDto() *SongAudiosDto {
	return poolSongAudiosDto.Get().(*SongAudiosDto)
}

// ReleaseSongAudiosDto 释放SongAudiosDto
func ReleaseSongAudiosDto(v *SongAudiosDto) {
	v.Audios = v.Audios[:0]
	v.SongId = 0
	v.RefrainEndTime = 0
	v.RefrainStartTime = 0
	v.VideoPartStartTime = 0
	v.VideoPartEndTime = 0
	poolSongAudiosDto.Put(v)
}
