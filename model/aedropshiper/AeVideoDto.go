package aedropshiper

import (
	"sync"
)

// AeVideoDto 结构体
type AeVideoDto struct {
	// Video status
	MediaStatus string `json:"media_status,omitempty" xml:"media_status,omitempty"`
	// Type of video
	MediaType string `json:"media_type,omitempty" xml:"media_type,omitempty"`
	// The URL of the video cover image
	PosterUrl string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
	// Seller&#39;s master account ID
	AliMemberId int64 `json:"ali_member_id,omitempty" xml:"ali_member_id,omitempty"`
	// Video ID
	MediaId int64 `json:"media_id,omitempty" xml:"media_id,omitempty"`
}

var poolAeVideoDto = sync.Pool{
	New: func() any {
		return new(AeVideoDto)
	},
}

// GetAeVideoDto() 从对象池中获取AeVideoDto
func GetAeVideoDto() *AeVideoDto {
	return poolAeVideoDto.Get().(*AeVideoDto)
}

// ReleaseAeVideoDto 释放AeVideoDto
func ReleaseAeVideoDto(v *AeVideoDto) {
	v.MediaStatus = ""
	v.MediaType = ""
	v.PosterUrl = ""
	v.AliMemberId = 0
	v.MediaId = 0
	poolAeVideoDto.Put(v)
}
