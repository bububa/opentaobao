package idle

import (
	"sync"
)

// ItemSkuDto 结构体
type ItemSkuDto struct {
	// sku属性
	PropList []ItemPvPairDto `json:"prop_list,omitempty" xml:"prop_list>item_pv_pair_dto,omitempty"`
	// sku价格，单位分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// sku库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItemSkuDto = sync.Pool{
	New: func() any {
		return new(ItemSkuDto)
	},
}

// GetItemSkuDto() 从对象池中获取ItemSkuDto
func GetItemSkuDto() *ItemSkuDto {
	return poolItemSkuDto.Get().(*ItemSkuDto)
}

// ReleaseItemSkuDto 释放ItemSkuDto
func ReleaseItemSkuDto(v *ItemSkuDto) {
	v.PropList = v.PropList[:0]
	v.Price = 0
	v.Quantity = 0
	v.SkuId = 0
	v.ItemId = 0
	poolItemSkuDto.Put(v)
}
