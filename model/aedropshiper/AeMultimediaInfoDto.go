package aedropshiper

import (
	"sync"
)

// AeMultimediaInfoDto 结构体
type AeMultimediaInfoDto struct {
	// Video information
	AeVideoDtos []AeVideoDto `json:"ae_video_dtos,omitempty" xml:"ae_video_dtos>ae_video_dto,omitempty"`
	// List of main images of the product
	ImageUrls string `json:"image_urls,omitempty" xml:"image_urls,omitempty"`
}

var poolAeMultimediaInfoDto = sync.Pool{
	New: func() any {
		return new(AeMultimediaInfoDto)
	},
}

// GetAeMultimediaInfoDto() 从对象池中获取AeMultimediaInfoDto
func GetAeMultimediaInfoDto() *AeMultimediaInfoDto {
	return poolAeMultimediaInfoDto.Get().(*AeMultimediaInfoDto)
}

// ReleaseAeMultimediaInfoDto 释放AeMultimediaInfoDto
func ReleaseAeMultimediaInfoDto(v *AeMultimediaInfoDto) {
	v.AeVideoDtos = v.AeVideoDtos[:0]
	v.ImageUrls = ""
	poolAeMultimediaInfoDto.Put(v)
}
