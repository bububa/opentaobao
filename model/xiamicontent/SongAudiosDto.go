package xiamicontent

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
}
