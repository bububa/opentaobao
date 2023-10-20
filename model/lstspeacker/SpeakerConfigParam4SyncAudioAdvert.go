package lstspeacker

// SpeakerConfigParam4syncAudioAdvert 结构体
type SpeakerConfigParam4syncAudioAdvert struct {
	// 广告列表
	AdvertList []SpeakerConfigAudioAdvert `json:"advert_list,omitempty" xml:"advert_list>speaker_config_audio_advert,omitempty"`
}
