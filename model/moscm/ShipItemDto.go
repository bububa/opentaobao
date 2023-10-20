package moscm

import (
	"sync"
)

// ShipItemDto 结构体
type ShipItemDto struct {
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品sku
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 唯一标识
	Id string `json:"id,omitempty" xml:"id,omitempty"`
}

var poolShipItemDto = sync.Pool{
	New: func() any {
		return new(ShipItemDto)
	},
}

// GetShipItemDto() 从对象池中获取ShipItemDto
func GetShipItemDto() *ShipItemDto {
	return poolShipItemDto.Get().(*ShipItemDto)
}

// ReleaseShipItemDto 释放ShipItemDto
func ReleaseShipItemDto(v *ShipItemDto) {
	v.Quantity = ""
	v.SkuId = ""
	v.Id = ""
	poolShipItemDto.Put(v)
}
