package lstvending

// VendingShippedEquipmentDto 结构体
type VendingShippedEquipmentDto struct {
	// 供应商设备唯一编码
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 设备型号清单ID
	OrderItemId int64 `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
}
