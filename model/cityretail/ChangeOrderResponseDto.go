package cityretail

import (
	"sync"
)

// ChangeOrderResponseDto 结构体
type ChangeOrderResponseDto struct {
	// 淘宝订单id
	TbOrderId string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
	// 转仓后的新仓
	NewWarehouseCode string `json:"new_warehouse_code,omitempty" xml:"new_warehouse_code,omitempty"`
	// 转仓前的原仓
	OriginWarehouseCode string `json:"origin_warehouse_code,omitempty" xml:"origin_warehouse_code,omitempty"`
	// 履约单id
	FulfillOrderId string `json:"fulfill_order_id,omitempty" xml:"fulfill_order_id,omitempty"`
	// 取货码
	PickupCode string `json:"pickup_code,omitempty" xml:"pickup_code,omitempty"`
}

var poolChangeOrderResponseDto = sync.Pool{
	New: func() any {
		return new(ChangeOrderResponseDto)
	},
}

// GetChangeOrderResponseDto() 从对象池中获取ChangeOrderResponseDto
func GetChangeOrderResponseDto() *ChangeOrderResponseDto {
	return poolChangeOrderResponseDto.Get().(*ChangeOrderResponseDto)
}

// ReleaseChangeOrderResponseDto 释放ChangeOrderResponseDto
func ReleaseChangeOrderResponseDto(v *ChangeOrderResponseDto) {
	v.TbOrderId = ""
	v.NewWarehouseCode = ""
	v.OriginWarehouseCode = ""
	v.FulfillOrderId = ""
	v.PickupCode = ""
	poolChangeOrderResponseDto.Put(v)
}
