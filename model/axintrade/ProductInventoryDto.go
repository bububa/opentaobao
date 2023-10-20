package axintrade

import (
	"sync"
)

// ProductInventoryDto 结构体
type ProductInventoryDto struct {
	// 日期 格式yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 库存数量
	InvCount int64 `json:"inv_count,omitempty" xml:"inv_count,omitempty"`
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
	v.Date = ""
	v.InvCount = 0
	poolProductInventoryDto.Put(v)
}
