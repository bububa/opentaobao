package waybill

import (
	"sync"
)

// WarehouseDto 结构体
type WarehouseDto struct {
	// 仓名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 仓id
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
}

var poolWarehouseDto = sync.Pool{
	New: func() any {
		return new(WarehouseDto)
	},
}

// GetWarehouseDto() 从对象池中获取WarehouseDto
func GetWarehouseDto() *WarehouseDto {
	return poolWarehouseDto.Get().(*WarehouseDto)
}

// ReleaseWarehouseDto 释放WarehouseDto
func ReleaseWarehouseDto(v *WarehouseDto) {
	v.WarehouseName = ""
	v.WarehouseId = 0
	poolWarehouseDto.Put(v)
}
