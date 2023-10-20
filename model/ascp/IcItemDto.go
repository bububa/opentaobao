package ascp

import (
	"sync"
)

// IcItemDto 结构体
type IcItemDto struct {
	// 店铺宝贝ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 店铺宝贝skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolIcItemDto = sync.Pool{
	New: func() any {
		return new(IcItemDto)
	},
}

// GetIcItemDto() 从对象池中获取IcItemDto
func GetIcItemDto() *IcItemDto {
	return poolIcItemDto.Get().(*IcItemDto)
}

// ReleaseIcItemDto 释放IcItemDto
func ReleaseIcItemDto(v *IcItemDto) {
	v.ItemId = ""
	v.SkuId = ""
	poolIcItemDto.Put(v)
}
