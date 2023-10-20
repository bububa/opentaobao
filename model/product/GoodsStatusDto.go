package product

import (
	"sync"
)

// GoodsStatusDto 结构体
type GoodsStatusDto struct {
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolGoodsStatusDto = sync.Pool{
	New: func() any {
		return new(GoodsStatusDto)
	},
}

// GetGoodsStatusDto() 从对象池中获取GoodsStatusDto
func GetGoodsStatusDto() *GoodsStatusDto {
	return poolGoodsStatusDto.Get().(*GoodsStatusDto)
}

// ReleaseGoodsStatusDto 释放GoodsStatusDto
func ReleaseGoodsStatusDto(v *GoodsStatusDto) {
	v.Price = ""
	v.Status = 0
	v.GoodsId = 0
	poolGoodsStatusDto.Put(v)
}
