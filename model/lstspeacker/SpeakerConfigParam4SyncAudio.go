package lstspeacker

// SpeakerConfigParam4SyncAudio 结构体
type SpeakerConfigParam4SyncAudio struct {
	// 音频集信息
	AudioInfo *SpeakerConfigAudioInfo `json:"audio_info,omitempty" xml:"audio_info,omitempty"`
	// 音频列表
	AudioList []SpeakerConfigAudioList `json:"audio_list,omitempty" xml:"audio_list>speaker_config_audio_list,omitempty"`
}
