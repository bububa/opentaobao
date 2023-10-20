package icbu

import (
	"sync"
)

// ProductInventoryDto 结构体
type ProductInventoryDto struct {
	// 库存编码，为空时表示默认国内仓
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 库存值
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
}

var poolProductInventoryDto = sync.Pool{
	New: func() any {
		return new(ProductInventoryDto)
	},
}

// GetProductInventoryDto() 从对象池中获取ProductInventoryDto
func GetProductInventoryDto() *ProductInventoryDto {
	return poolProductInventoryDto.Get().(*ProductInventoryDto)
}

// ReleaseProductInventoryDto 释放ProductInventoryDto
func ReleaseProductInventoryDto(v *ProductInventoryDto) {
	v.StoreCode = ""
	v.Inventory = 0
	poolProductInventoryDto.Put(v)
}
