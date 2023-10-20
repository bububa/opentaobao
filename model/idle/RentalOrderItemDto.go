package idle

import (
	"sync"
)

// RentalOrderItemDto 结构体
type RentalOrderItemDto struct {
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sku信息
	SkuDTO *ItemSkuDto `json:"sku_d_t_o,omitempty" xml:"sku_d_t_o,omitempty"`
	// 价格信息
	Price *PriceDto `json:"price,omitempty" xml:"price,omitempty"`
}

var poolRentalOrderItemDto = sync.Pool{
	New: func() any {
		return new(RentalOrderItemDto)
	},
}

// GetRentalOrderItemDto() 从对象池中获取RentalOrderItemDto
func GetRentalOrderItemDto() *RentalOrderItemDto {
	return poolRentalOrderItemDto.Get().(*RentalOrderItemDto)
}

// ReleaseRentalOrderItemDto 释放RentalOrderItemDto
func ReleaseRentalOrderItemDto(v *RentalOrderItemDto) {
	v.ItemId = 0
	v.SkuDTO = nil
	v.Price = nil
	poolRentalOrderItemDto.Put(v)
}
