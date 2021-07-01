package lstspeacker

// SpeakerConfigParam4SyncAudioAdvert 结构体
type SpeakerConfigParam4SyncAudioAdvert struct {
	// 广告列表
	AdvertList []SpeakerConfigAudioAdvert `json:"advert_list,omitempty" xml:"advert_list>speaker_config_audio_advert,omitempty"`
}
