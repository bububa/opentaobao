package xiamicontent

// AudioDto 结构体
type AudioDto struct {
	// 试听文件地址
	ListenUrl string `json:"listen_url,omitempty" xml:"listen_url,omitempty"`
	// 格式
	Format string `json:"format,omitempty" xml:"format,omitempty"`
	// 音频封面
	AudioCover string `json:"audio_cover,omitempty" xml:"audio_cover,omitempty"`
	// 音频名称
	AudioName string `json:"audio_name,omitempty" xml:"audio_name,omitempty"`
	// 音频描述
	AudioDesc string `json:"audio_desc,omitempty" xml:"audio_desc,omitempty"`
	// 时长
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 码率
	Rate int64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// 文件大小
	FileSize int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
	// 音频id
	AudioId int64 `json:"audio_id,omitempty" xml:"audio_id,omitempty"`
	// 位数
	Bits int64 `json:"bits,omitempty" xml:"bits,omitempty"`
	// 采样率
	SampleRate int64 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
	// 音质
	Quality int64 `json:"quality,omitempty" xml:"quality,omitempty"`
	// 超时时间
	Expire int64 `json:"expire,omitempty" xml:"expire,omitempty"`
}
