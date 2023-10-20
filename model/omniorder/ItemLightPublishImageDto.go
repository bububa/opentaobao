package omniorder

import (
	"sync"
)

// ItemLightPublishImageDto 结构体
type ItemLightPublishImageDto struct {
	// 图片url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolItemLightPublishImageDto = sync.Pool{
	New: func() any {
		return new(ItemLightPublishImageDto)
	},
}

// GetItemLightPublishImageDto() 从对象池中获取ItemLightPublishImageDto
func GetItemLightPublishImageDto() *ItemLightPublishImageDto {
	return poolItemLightPublishImageDto.Get().(*ItemLightPublishImageDto)
}

// ReleaseItemLightPublishImageDto 释放ItemLightPublishImageDto
func ReleaseItemLightPublishImageDto(v *ItemLightPublishImageDto) {
	v.Url = ""
	poolItemLightPublishImageDto.Put(v)
}
