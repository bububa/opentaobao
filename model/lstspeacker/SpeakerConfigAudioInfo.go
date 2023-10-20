package lstspeacker

import (
	"sync"
)

// SpeakerConfigAudioInfo 结构体
type SpeakerConfigAudioInfo struct {
	// 音频集名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 音频集开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 音频集结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 描述
	Describes string `json:"describes,omitempty" xml:"describes,omitempty"`
	// 音频封面
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolSpeakerConfigAudioInfo = sync.Pool{
	New: func() any {
		return new(SpeakerConfigAudioInfo)
	},
}

// GetSpeakerConfigAudioInfo() 从对象池中获取SpeakerConfigAudioInfo
func GetSpeakerConfigAudioInfo() *SpeakerConfigAudioInfo {
	return poolSpeakerConfigAudioInfo.Get().(*SpeakerConfigAudioInfo)
}

// ReleaseSpeakerConfigAudioInfo 释放SpeakerConfigAudioInfo
func ReleaseSpeakerConfigAudioInfo(v *SpeakerConfigAudioInfo) {
	v.Name = ""
	v.StartTime = ""
	v.EndTime = ""
	v.Describes = ""
	v.Url = ""
	poolSpeakerConfigAudioInfo.Put(v)
}
