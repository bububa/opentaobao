package tmallnr

import (
	"sync"
)

// StoreInvetoryDto 结构体
type StoreInvetoryDto struct {
	// 商家的外部商品编码，写入值。
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 单据流水号，用于幂等操作
	BillNum string `json:"bill_num,omitempty" xml:"bill_num,omitempty"`
	// CERTAINTY 表示确定性库存
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// sku号
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 天猫商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 天猫后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolStoreInvetoryDto = sync.Pool{
	New: func() any {
		return new(StoreInvetoryDto)
	},
}

// GetStoreInvetoryDto() 从对象池中获取StoreInvetoryDto
func GetStoreInvetoryDto() *StoreInvetoryDto {
	return poolStoreInvetoryDto.Get().(*StoreInvetoryDto)
}

// ReleaseStoreInvetoryDto 释放StoreInvetoryDto
func ReleaseStoreInvetoryDto(v *StoreInvetoryDto) {
	v.OuterId = ""
	v.BillNum = ""
	v.InventoryType = ""
	v.SkuId = 0
	v.ItemId = 0
	v.ScItemId = 0
	v.Quantity = 0
	poolStoreInvetoryDto.Put(v)
}
