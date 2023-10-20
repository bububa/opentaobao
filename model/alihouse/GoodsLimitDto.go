package alihouse

import (
	"sync"
)

// GoodsLimitDto 结构体
type GoodsLimitDto struct {
	// 商品级别限购: 一个商品ID, 同一用户限制能下单次数
	SingleItemLimit int64 `json:"single_item_limit,omitempty" xml:"single_item_limit,omitempty"`
	// SKU级别限购: 一个商品SKUid, 同一用户限制能下单次数
	SingleSkuLimit int64 `json:"single_sku_limit,omitempty" xml:"single_sku_limit,omitempty"`
}

var poolGoodsLimitDto = sync.Pool{
	New: func() any {
		return new(GoodsLimitDto)
	},
}

// GetGoodsLimitDto() 从对象池中获取GoodsLimitDto
func GetGoodsLimitDto() *GoodsLimitDto {
	return poolGoodsLimitDto.Get().(*GoodsLimitDto)
}

// ReleaseGoodsLimitDto 释放GoodsLimitDto
func ReleaseGoodsLimitDto(v *GoodsLimitDto) {
	v.SingleItemLimit = 0
	v.SingleSkuLimit = 0
	poolGoodsLimitDto.Put(v)
}
