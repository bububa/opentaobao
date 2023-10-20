package lstspeacker

import (
	"sync"
)

// SpeakerConfigParam4SyncAudioAdvert 结构体
type SpeakerConfigParam4SyncAudioAdvert struct {
	// 广告列表
	AdvertList []SpeakerConfigAudioAdvert `json:"advert_list,omitempty" xml:"advert_list>speaker_config_audio_advert,omitempty"`
}

var poolSpeakerConfigParam4SyncAudioAdvert = sync.Pool{
	New: func() any {
		return new(SpeakerConfigParam4SyncAudioAdvert)
	},
}

// GetSpeakerConfigParam4SyncAudioAdvert() 从对象池中获取SpeakerConfigParam4SyncAudioAdvert
func GetSpeakerConfigParam4SyncAudioAdvert() *SpeakerConfigParam4SyncAudioAdvert {
	return poolSpeakerConfigParam4SyncAudioAdvert.Get().(*SpeakerConfigParam4SyncAudioAdvert)
}

// ReleaseSpeakerConfigParam4SyncAudioAdvert 释放SpeakerConfigParam4SyncAudioAdvert
func ReleaseSpeakerConfigParam4SyncAudioAdvert(v *SpeakerConfigParam4SyncAudioAdvert) {
	v.AdvertList = v.AdvertList[:0]
	poolSpeakerConfigParam4SyncAudioAdvert.Put(v)
}
