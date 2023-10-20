package ott

import (
	"sync"
)

// ImageDto 结构体
type ImageDto struct {
	// imageUrl
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// type
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolImageDto = sync.Pool{
	New: func() any {
		return new(ImageDto)
	},
}

// GetImageDto() 从对象池中获取ImageDto
func GetImageDto() *ImageDto {
	return poolImageDto.Get().(*ImageDto)
}

// ReleaseImageDto 释放ImageDto
func ReleaseImageDto(v *ImageDto) {
	v.ImageUrl = ""
	v.Type = ""
	poolImageDto.Put(v)
}
