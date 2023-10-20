package lstspeacker

import (
	"sync"
)

// SpeakerFileDto 结构体
type SpeakerFileDto struct {
	// md5
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// fileId
	FileId string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

var poolSpeakerFileDto = sync.Pool{
	New: func() any {
		return new(SpeakerFileDto)
	},
}

// GetSpeakerFileDto() 从对象池中获取SpeakerFileDto
func GetSpeakerFileDto() *SpeakerFileDto {
	return poolSpeakerFileDto.Get().(*SpeakerFileDto)
}

// ReleaseSpeakerFileDto 释放SpeakerFileDto
func ReleaseSpeakerFileDto(v *SpeakerFileDto) {
	v.Md5 = ""
	v.Url = ""
	v.FileId = ""
	poolSpeakerFileDto.Put(v)
}
