package lstvending

import (
	"sync"
)

// VendingShippedEquipmentDto 结构体
type VendingShippedEquipmentDto struct {
	// 供应商设备唯一编码
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 设备型号清单ID
	OrderItemId int64 `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
}

var poolVendingShippedEquipmentDto = sync.Pool{
	New: func() any {
		return new(VendingShippedEquipmentDto)
	},
}

// GetVendingShippedEquipmentDto() 从对象池中获取VendingShippedEquipmentDto
func GetVendingShippedEquipmentDto() *VendingShippedEquipmentDto {
	return poolVendingShippedEquipmentDto.Get().(*VendingShippedEquipmentDto)
}

// ReleaseVendingShippedEquipmentDto 释放VendingShippedEquipmentDto
func ReleaseVendingShippedEquipmentDto(v *VendingShippedEquipmentDto) {
	v.EquipmentCode = ""
	v.OrderItemId = 0
	poolVendingShippedEquipmentDto.Put(v)
}
