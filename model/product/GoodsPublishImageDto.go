package product

import (
	"sync"
)

// GoodsPublishImageDto 结构体
type GoodsPublishImageDto struct {
	// 商品图片url
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolGoodsPublishImageDto = sync.Pool{
	New: func() any {
		return new(GoodsPublishImageDto)
	},
}

// GetGoodsPublishImageDto() 从对象池中获取GoodsPublishImageDto
func GetGoodsPublishImageDto() *GoodsPublishImageDto {
	return poolGoodsPublishImageDto.Get().(*GoodsPublishImageDto)
}

// ReleaseGoodsPublishImageDto 释放GoodsPublishImageDto
func ReleaseGoodsPublishImageDto(v *GoodsPublishImageDto) {
	v.ImageUrl = ""
	poolGoodsPublishImageDto.Put(v)
}
