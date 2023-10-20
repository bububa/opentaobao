package iot

import (
	"sync"
)

// ImageUrlDto 结构体
type ImageUrlDto struct {
	// 默认图片
	Img string `json:"img,omitempty" xml:"img,omitempty"`
	// 大图
	Large string `json:"large,omitempty" xml:"large,omitempty"`
	// 中图
	Medium string `json:"medium,omitempty" xml:"medium,omitempty"`
	// 小图
	Small string `json:"small,omitempty" xml:"small,omitempty"`
}

var poolImageUrlDto = sync.Pool{
	New: func() any {
		return new(ImageUrlDto)
	},
}

// GetImageUrlDto() 从对象池中获取ImageUrlDto
func GetImageUrlDto() *ImageUrlDto {
	return poolImageUrlDto.Get().(*ImageUrlDto)
}

// ReleaseImageUrlDto 释放ImageUrlDto
func ReleaseImageUrlDto(v *ImageUrlDto) {
	v.Img = ""
	v.Large = ""
	v.Medium = ""
	v.Small = ""
	poolImageUrlDto.Put(v)
}
