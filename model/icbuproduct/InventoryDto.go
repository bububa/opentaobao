package icbuproduct

import (
	"sync"
)

// InventoryDto 结构体
type InventoryDto struct {
	// 库存的仓编码,根据商品查询返回的仓编码进行设置,不同的客户类型,仓编码会不一样
	InventoryCode string `json:"inventory_code,omitempty" xml:"inventory_code,omitempty"`
	// 操作类型，加库存或者减库存,加库存:plus,减库存:sub
	Operate string `json:"operate,omitempty" xml:"operate,omitempty"`
	// 待更新库存的SKUid,如果没有skuId,设置成-1
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 库存变动值
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
}

var poolInventoryDto = sync.Pool{
	New: func() any {
		return new(InventoryDto)
	},
}

// GetInventoryDto() 从对象池中获取InventoryDto
func GetInventoryDto() *InventoryDto {
	return poolInventoryDto.Get().(*InventoryDto)
}

// ReleaseInventoryDto 释放InventoryDto
func ReleaseInventoryDto(v *InventoryDto) {
	v.InventoryCode = ""
	v.Operate = ""
	v.SkuId = 0
	v.Inventory = 0
	poolInventoryDto.Put(v)
}
