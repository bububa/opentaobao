package idle

import (
	"sync"
)

// MediaDto 结构体
type MediaDto struct {
	// 商品主图列表
	Images []ImageInfoDto `json:"images,omitempty" xml:"images>image_info_dto,omitempty"`
	// 商品详情图片列表
	PropImages []ImageInfoDto `json:"prop_images,omitempty" xml:"prop_images>image_info_dto,omitempty"`
}

var poolMediaDto = sync.Pool{
	New: func() any {
		return new(MediaDto)
	},
}

// GetMediaDto() 从对象池中获取MediaDto
func GetMediaDto() *MediaDto {
	return poolMediaDto.Get().(*MediaDto)
}

// ReleaseMediaDto 释放MediaDto
func ReleaseMediaDto(v *MediaDto) {
	v.Images = v.Images[:0]
	v.PropImages = v.PropImages[:0]
	poolMediaDto.Put(v)
}
