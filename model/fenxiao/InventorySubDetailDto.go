package fenxiao

import (
	"sync"
)

// InventorySubDetailDto 结构体
type InventorySubDetailDto struct {
	// ONLINE_INVENTORY:线上可售卖库存。SHARE_INVENTORY：线下独享库存，门店自提可用
	InvBizCode string `json:"inv_biz_code,omitempty" xml:"inv_biz_code,omitempty"`
	// 预扣库存数
	ReserveQuantity int64 `json:"reserve_quantity,omitempty" xml:"reserve_quantity,omitempty"`
	// 仓库物理库存数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 占用库存数
	OccupyQuantity int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
}

var poolInventorySubDetailDto = sync.Pool{
	New: func() any {
		return new(InventorySubDetailDto)
	},
}

// GetInventorySubDetailDto() 从对象池中获取InventorySubDetailDto
func GetInventorySubDetailDto() *InventorySubDetailDto {
	return poolInventorySubDetailDto.Get().(*InventorySubDetailDto)
}

// ReleaseInventorySubDetailDto 释放InventorySubDetailDto
func ReleaseInventorySubDetailDto(v *InventorySubDetailDto) {
	v.InvBizCode = ""
	v.ReserveQuantity = 0
	v.Quantity = 0
	v.OccupyQuantity = 0
	poolInventorySubDetailDto.Put(v)
}
