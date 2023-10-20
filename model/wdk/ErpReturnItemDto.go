package wdk

import (
	"sync"
)

// ErpReturnItemDto 结构体
type ErpReturnItemDto struct {
	// 数量
	Count string `json:"count,omitempty" xml:"count,omitempty"`
	// 库位号，退货库位号
	CabinetCode string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
	// 商品code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolErpReturnItemDto = sync.Pool{
	New: func() any {
		return new(ErpReturnItemDto)
	},
}

// GetErpReturnItemDto() 从对象池中获取ErpReturnItemDto
func GetErpReturnItemDto() *ErpReturnItemDto {
	return poolErpReturnItemDto.Get().(*ErpReturnItemDto)
}

// ReleaseErpReturnItemDto 释放ErpReturnItemDto
func ReleaseErpReturnItemDto(v *ErpReturnItemDto) {
	v.Count = ""
	v.CabinetCode = ""
	v.ItemCode = ""
	v.WarehouseCode = ""
	poolErpReturnItemDto.Put(v)
}
