package product

import (
	"sync"
)

// GoodsImageDto 结构体
type GoodsImageDto struct {
	// 原图url
	OriginImage string `json:"origin_image,omitempty" xml:"origin_image,omitempty"`
	// 备注
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 图片id
	ImageId int64 `json:"image_id,omitempty" xml:"image_id,omitempty"`
}

var poolGoodsImageDto = sync.Pool{
	New: func() any {
		return new(GoodsImageDto)
	},
}

// GetGoodsImageDto() 从对象池中获取GoodsImageDto
func GetGoodsImageDto() *GoodsImageDto {
	return poolGoodsImageDto.Get().(*GoodsImageDto)
}

// ReleaseGoodsImageDto 释放GoodsImageDto
func ReleaseGoodsImageDto(v *GoodsImageDto) {
	v.OriginImage = ""
	v.Note = ""
	v.ImageId = 0
	poolGoodsImageDto.Put(v)
}
