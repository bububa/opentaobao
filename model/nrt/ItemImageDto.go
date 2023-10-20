package nrt

import (
	"sync"
)

// ItemImageDto 结构体
type ItemImageDto struct {
	// 图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片次序
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}

var poolItemImageDto = sync.Pool{
	New: func() any {
		return new(ItemImageDto)
	},
}

// GetItemImageDto() 从对象池中获取ItemImageDto
func GetItemImageDto() *ItemImageDto {
	return poolItemImageDto.Get().(*ItemImageDto)
}

// ReleaseItemImageDto 释放ItemImageDto
func ReleaseItemImageDto(v *ItemImageDto) {
	v.Url = ""
	v.Position = 0
	poolItemImageDto.Put(v)
}
