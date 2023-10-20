package simba

import (
	"sync"
)

// VideoFeedDto 结构体
type VideoFeedDto struct {
	// 视频状态描述
	StateDesc string `json:"state_desc,omitempty" xml:"state_desc,omitempty"`
	// 视频比例
	Ratio string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// 视频类型
	SizeType int64 `json:"size_type,omitempty" xml:"size_type,omitempty"`
	// 视频ID
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 视频宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 视频状态码
	State int64 `json:"state,omitempty" xml:"state,omitempty"`
	// 视频高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}

var poolVideoFeedDto = sync.Pool{
	New: func() any {
		return new(VideoFeedDto)
	},
}

// GetVideoFeedDto() 从对象池中获取VideoFeedDto
func GetVideoFeedDto() *VideoFeedDto {
	return poolVideoFeedDto.Get().(*VideoFeedDto)
}

// ReleaseVideoFeedDto 释放VideoFeedDto
func ReleaseVideoFeedDto(v *VideoFeedDto) {
	v.StateDesc = ""
	v.Ratio = ""
	v.SizeType = 0
	v.VideoId = 0
	v.ItemId = 0
	v.Width = 0
	v.State = 0
	v.Height = 0
	poolVideoFeedDto.Put(v)
}
