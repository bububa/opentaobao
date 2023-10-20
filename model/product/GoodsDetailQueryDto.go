package product

import (
	"sync"
)

// GoodsDetailQueryDto 结构体
type GoodsDetailQueryDto struct {
	// 交易猫商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolGoodsDetailQueryDto = sync.Pool{
	New: func() any {
		return new(GoodsDetailQueryDto)
	},
}

// GetGoodsDetailQueryDto() 从对象池中获取GoodsDetailQueryDto
func GetGoodsDetailQueryDto() *GoodsDetailQueryDto {
	return poolGoodsDetailQueryDto.Get().(*GoodsDetailQueryDto)
}

// ReleaseGoodsDetailQueryDto 释放GoodsDetailQueryDto
func ReleaseGoodsDetailQueryDto(v *GoodsDetailQueryDto) {
	v.GoodsId = 0
	poolGoodsDetailQueryDto.Put(v)
}
