package alihouse

import (
	"sync"
)

// VideoDataDto 结构体
type VideoDataDto struct {
	// 视频下载链接
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 视频封面
	CoverImage string `json:"cover_image,omitempty" xml:"cover_image,omitempty"`
	// 外部视频id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolVideoDataDto = sync.Pool{
	New: func() any {
		return new(VideoDataDto)
	},
}

// GetVideoDataDto() 从对象池中获取VideoDataDto
func GetVideoDataDto() *VideoDataDto {
	return poolVideoDataDto.Get().(*VideoDataDto)
}

// ReleaseVideoDataDto 释放VideoDataDto
func ReleaseVideoDataDto(v *VideoDataDto) {
	v.VideoUrl = ""
	v.CoverImage = ""
	v.Id = 0
	poolVideoDataDto.Put(v)
}
