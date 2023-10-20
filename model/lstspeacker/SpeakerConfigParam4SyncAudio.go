package lstspeacker

import (
	"sync"
)

// SpeakerConfigParam4SyncAudio 结构体
type SpeakerConfigParam4SyncAudio struct {
	// 音频列表
	AudioList []SpeakerConfigAudioList `json:"audio_list,omitempty" xml:"audio_list>speaker_config_audio_list,omitempty"`
	// 音频集信息
	AudioInfo *SpeakerConfigAudioInfo `json:"audio_info,omitempty" xml:"audio_info,omitempty"`
}

var poolSpeakerConfigParam4SyncAudio = sync.Pool{
	New: func() any {
		return new(SpeakerConfigParam4SyncAudio)
	},
}

// GetSpeakerConfigParam4SyncAudio() 从对象池中获取SpeakerConfigParam4SyncAudio
func GetSpeakerConfigParam4SyncAudio() *SpeakerConfigParam4SyncAudio {
	return poolSpeakerConfigParam4SyncAudio.Get().(*SpeakerConfigParam4SyncAudio)
}

// ReleaseSpeakerConfigParam4SyncAudio 释放SpeakerConfigParam4SyncAudio
func ReleaseSpeakerConfigParam4SyncAudio(v *SpeakerConfigParam4SyncAudio) {
	v.AudioList = v.AudioList[:0]
	v.AudioInfo = nil
	poolSpeakerConfigParam4SyncAudio.Put(v)
}
