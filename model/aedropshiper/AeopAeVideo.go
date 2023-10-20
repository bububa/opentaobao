package aedropshiper

import (
	"sync"
)

// AeopAeVideo 结构体
type AeopAeVideo struct {
	// 视频封面图片的URL
	PosterUrl string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
	// 视频的类型
	MediaType string `json:"media_type,omitempty" xml:"media_type,omitempty"`
	// 视频的状态
	MediaStatus string `json:"media_status,omitempty" xml:"media_status,omitempty"`
	// 视频ID
	MediaId int64 `json:"media_id,omitempty" xml:"media_id,omitempty"`
	// 卖家主账户ID
	AliMemberId int64 `json:"ali_member_id,omitempty" xml:"ali_member_id,omitempty"`
}

var poolAeopAeVideo = sync.Pool{
	New: func() any {
		return new(AeopAeVideo)
	},
}

// GetAeopAeVideo() 从对象池中获取AeopAeVideo
func GetAeopAeVideo() *AeopAeVideo {
	return poolAeopAeVideo.Get().(*AeopAeVideo)
}

// ReleaseAeopAeVideo 释放AeopAeVideo
func ReleaseAeopAeVideo(v *AeopAeVideo) {
	v.PosterUrl = ""
	v.MediaType = ""
	v.MediaStatus = ""
	v.MediaId = 0
	v.AliMemberId = 0
	poolAeopAeVideo.Put(v)
}
