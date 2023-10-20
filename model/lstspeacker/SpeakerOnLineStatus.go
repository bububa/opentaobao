package lstspeacker

import (
	"sync"
)

// SpeakerOnLineStatus 结构体
type SpeakerOnLineStatus struct {
	// 状态吗
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 状态名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolSpeakerOnLineStatus = sync.Pool{
	New: func() any {
		return new(SpeakerOnLineStatus)
	},
}

// GetSpeakerOnLineStatus() 从对象池中获取SpeakerOnLineStatus
func GetSpeakerOnLineStatus() *SpeakerOnLineStatus {
	return poolSpeakerOnLineStatus.Get().(*SpeakerOnLineStatus)
}

// ReleaseSpeakerOnLineStatus 释放SpeakerOnLineStatus
func ReleaseSpeakerOnLineStatus(v *SpeakerOnLineStatus) {
	v.Code = ""
	v.Name = ""
	poolSpeakerOnLineStatus.Put(v)
}
