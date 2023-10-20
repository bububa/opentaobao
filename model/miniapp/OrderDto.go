package miniapp

import (
	"sync"
)

// OrderDto 结构体
type OrderDto struct {
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderDto = sync.Pool{
	New: func() any {
		return new(OrderDto)
	},
}

// GetOrderDto() 从对象池中获取OrderDto
func GetOrderDto() *OrderDto {
	return poolOrderDto.Get().(*OrderDto)
}

// ReleaseOrderDto 释放OrderDto
func ReleaseOrderDto(v *OrderDto) {
	v.OrderId = ""
	poolOrderDto.Put(v)
}
