package icbulogistics

import (
	"sync"
)

// OrderDetailDto 结构体
type OrderDetailDto struct {
	// 条码Base64
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 物流订单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 仓库信息
	Warehouse *WarehouseDto `json:"warehouse,omitempty" xml:"warehouse,omitempty"`
	// 上门揽收信息
	PickupInfo *PickupInfoDto `json:"pickup_info,omitempty" xml:"pickup_info,omitempty"`
}

var poolOrderDetailDto = sync.Pool{
	New: func() any {
		return new(OrderDetailDto)
	},
}

// GetOrderDetailDto() 从对象池中获取OrderDetailDto
func GetOrderDetailDto() *OrderDetailDto {
	return poolOrderDetailDto.Get().(*OrderDetailDto)
}

// ReleaseOrderDetailDto 释放OrderDetailDto
func ReleaseOrderDetailDto(v *OrderDetailDto) {
	v.BarCode = ""
	v.OrderNumber = ""
	v.Warehouse = nil
	v.PickupInfo = nil
	poolOrderDetailDto.Put(v)
}
