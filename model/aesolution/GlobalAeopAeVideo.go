package aesolution

import (
	"sync"
)

// GlobalAeopAeVideo 结构体
type GlobalAeopAeVideo struct {
	// The status of the video
	MediaStatus string `json:"media_status,omitempty" xml:"media_status,omitempty"`
	// The type of video
	MediaType string `json:"media_type,omitempty" xml:"media_type,omitempty"`
	// URL of the video cover image
	PosterUrl string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
	// Seller admin ID
	AliMemberId int64 `json:"ali_member_id,omitempty" xml:"ali_member_id,omitempty"`
	// Video ID
	MediaId int64 `json:"media_id,omitempty" xml:"media_id,omitempty"`
}

var poolGlobalAeopAeVideo = sync.Pool{
	New: func() any {
		return new(GlobalAeopAeVideo)
	},
}

// GetGlobalAeopAeVideo() 从对象池中获取GlobalAeopAeVideo
func GetGlobalAeopAeVideo() *GlobalAeopAeVideo {
	return poolGlobalAeopAeVideo.Get().(*GlobalAeopAeVideo)
}

// ReleaseGlobalAeopAeVideo 释放GlobalAeopAeVideo
func ReleaseGlobalAeopAeVideo(v *GlobalAeopAeVideo) {
	v.MediaStatus = ""
	v.MediaType = ""
	v.PosterUrl = ""
	v.AliMemberId = 0
	v.MediaId = 0
	poolGlobalAeopAeVideo.Put(v)
}
