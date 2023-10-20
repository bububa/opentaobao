package tmallnr

import (
	"sync"
)

// StoreDto 结构体
type StoreDto struct {
	// 门店商品，最大列表长度：20
	StoreInventories []StoreInvetoryDto `json:"store_inventories,omitempty" xml:"store_inventories>store_invetory_dto,omitempty"`
	// 库存来源的类型；STORE表示门店
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 门店对应的storeID值
	WarehouseId string `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
}

var poolStoreDto = sync.Pool{
	New: func() any {
		return new(StoreDto)
	},
}

// GetStoreDto() 从对象池中获取StoreDto
func GetStoreDto() *StoreDto {
	return poolStoreDto.Get().(*StoreDto)
}

// ReleaseStoreDto 释放StoreDto
func ReleaseStoreDto(v *StoreDto) {
	v.StoreInventories = v.StoreInventories[:0]
	v.WarehouseType = ""
	v.WarehouseId = ""
	poolStoreDto.Put(v)
}
