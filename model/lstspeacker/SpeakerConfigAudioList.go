package lstspeacker

// SpeakerConfigAudioList 
type SpeakerConfigAudioList struct {
    // 音频名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 音频播放时间秒
    Length   int64 `json:"length,omitempty" xml:"length,omitempty"`
    // 音频播放地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 音频md5
    Md5   string `json:"md5,omitempty" xml:"md5,omitempty"`
    // 音频播放顺序
    PlayPos   int64 `json:"play_pos,omitempty" xml:"play_pos,omitempty"`
    // 音频播放开始时间
    PlayStartDateTime   int64 `json:"play_start_date_time,omitempty" xml:"play_start_date_time,omitempty"`
    // 音频播放结束时间
    PlayEndDateTime   int64 `json:"play_end_date_time,omitempty" xml:"play_end_date_time,omitempty"`
}
