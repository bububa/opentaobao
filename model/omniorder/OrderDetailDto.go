package omniorder

import (
	"sync"
)

// OrderDetailDto 结构体
type OrderDetailDto struct {
	// 取件信息
	PickUpInfos []Content `json:"pick_up_infos,omitempty" xml:"pick_up_infos>content,omitempty"`
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
	v.PickUpInfos = v.PickUpInfos[:0]
	poolOrderDetailDto.Put(v)
}
