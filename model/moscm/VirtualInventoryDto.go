package moscm

import (
	"sync"
)

// VirtualInventoryDto 结构体
type VirtualInventoryDto struct {
	// 可售库存数量
	AvailableQuantity string `json:"available_quantity,omitempty" xml:"available_quantity,omitempty"`
	// 银泰专柜号
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 占库数量
	OccupyQty string `json:"occupy_qty,omitempty" xml:"occupy_qty,omitempty"`
	// 外部专柜号
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 外部商品id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 在库数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolVirtualInventoryDto = sync.Pool{
	New: func() any {
		return new(VirtualInventoryDto)
	},
}

// GetVirtualInventoryDto() 从对象池中获取VirtualInventoryDto
func GetVirtualInventoryDto() *VirtualInventoryDto {
	return poolVirtualInventoryDto.Get().(*VirtualInventoryDto)
}

// ReleaseVirtualInventoryDto 释放VirtualInventoryDto
func ReleaseVirtualInventoryDto(v *VirtualInventoryDto) {
	v.AvailableQuantity = ""
	v.CounterId = ""
	v.OccupyQty = ""
	v.OutCounterId = ""
	v.OutId = ""
	v.Quantity = ""
	poolVirtualInventoryDto.Put(v)
}
