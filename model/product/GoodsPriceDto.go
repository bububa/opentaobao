package product

import (
	"sync"
)

// GoodsPriceDto 结构体
type GoodsPriceDto struct {
	// 商品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 外部商品id对象
	ExternalGoodsId *ExternalGoodsIdDto `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
}

var poolGoodsPriceDto = sync.Pool{
	New: func() any {
		return new(GoodsPriceDto)
	},
}

// GetGoodsPriceDto() 从对象池中获取GoodsPriceDto
func GetGoodsPriceDto() *GoodsPriceDto {
	return poolGoodsPriceDto.Get().(*GoodsPriceDto)
}

// ReleaseGoodsPriceDto 释放GoodsPriceDto
func ReleaseGoodsPriceDto(v *GoodsPriceDto) {
	v.Price = ""
	v.ExternalGoodsId = nil
	poolGoodsPriceDto.Put(v)
}
