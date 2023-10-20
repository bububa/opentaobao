package inventory

import (
	"sync"
)

// InventorySyncDto 结构体
type InventorySyncDto struct {
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 库存数量
	ItemAmount int64 `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
}

var poolInventorySyncDto = sync.Pool{
	New: func() any {
		return new(InventorySyncDto)
	},
}

// GetInventorySyncDto() 从对象池中获取InventorySyncDto
func GetInventorySyncDto() *InventorySyncDto {
	return poolInventorySyncDto.Get().(*InventorySyncDto)
}

// ReleaseInventorySyncDto 释放InventorySyncDto
func ReleaseInventorySyncDto(v *InventorySyncDto) {
	v.ItemId = 0
	v.ItemAmount = 0
	poolInventorySyncDto.Put(v)
}
