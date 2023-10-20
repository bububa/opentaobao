package idle

import (
	"sync"
)

// ImageInfoDto 结构体
type ImageInfoDto struct {
	// 图片id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolImageInfoDto = sync.Pool{
	New: func() any {
		return new(ImageInfoDto)
	},
}

// GetImageInfoDto() 从对象池中获取ImageInfoDto
func GetImageInfoDto() *ImageInfoDto {
	return poolImageInfoDto.Get().(*ImageInfoDto)
}

// ReleaseImageInfoDto 释放ImageInfoDto
func ReleaseImageInfoDto(v *ImageInfoDto) {
	v.Id = 0
	poolImageInfoDto.Put(v)
}
