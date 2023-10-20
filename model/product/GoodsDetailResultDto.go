package product

import (
	"sync"
)

// GoodsDetailResultDto 结构体
type GoodsDetailResultDto struct {
	// 商品详情信息
	GoodsDetail *ExternalGoodsDetailDto `json:"goods_detail,omitempty" xml:"goods_detail,omitempty"`
	// 商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolGoodsDetailResultDto = sync.Pool{
	New: func() any {
		return new(GoodsDetailResultDto)
	},
}

// GetGoodsDetailResultDto() 从对象池中获取GoodsDetailResultDto
func GetGoodsDetailResultDto() *GoodsDetailResultDto {
	return poolGoodsDetailResultDto.Get().(*GoodsDetailResultDto)
}

// ReleaseGoodsDetailResultDto 释放GoodsDetailResultDto
func ReleaseGoodsDetailResultDto(v *GoodsDetailResultDto) {
	v.GoodsDetail = nil
	v.GoodsId = 0
	poolGoodsDetailResultDto.Put(v)
}
