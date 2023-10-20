package alihealthmedical

import (
	"sync"
)

// OrderQueryRequestDto 结构体
type OrderQueryRequestDto struct {
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 互联网医院编码
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}

var poolOrderQueryRequestDto = sync.Pool{
	New: func() any {
		return new(OrderQueryRequestDto)
	},
}

// GetOrderQueryRequestDto() 从对象池中获取OrderQueryRequestDto
func GetOrderQueryRequestDto() *OrderQueryRequestDto {
	return poolOrderQueryRequestDto.Get().(*OrderQueryRequestDto)
}

// ReleaseOrderQueryRequestDto 释放OrderQueryRequestDto
func ReleaseOrderQueryRequestDto(v *OrderQueryRequestDto) {
	v.OrderId = ""
	v.HospitalId = ""
	poolOrderQueryRequestDto.Put(v)
}
