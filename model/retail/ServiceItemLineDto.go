package retail

import (
	"sync"
)

// ServiceItemLineDto 结构体
type ServiceItemLineDto struct {
	// 服务商品编码
	ServiceSpuCode string `json:"service_spu_code,omitempty" xml:"service_spu_code,omitempty"`
	// 服务子订单id
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolServiceItemLineDto = sync.Pool{
	New: func() any {
		return new(ServiceItemLineDto)
	},
}

// GetServiceItemLineDto() 从对象池中获取ServiceItemLineDto
func GetServiceItemLineDto() *ServiceItemLineDto {
	return poolServiceItemLineDto.Get().(*ServiceItemLineDto)
}

// ReleaseServiceItemLineDto 释放ServiceItemLineDto
func ReleaseServiceItemLineDto(v *ServiceItemLineDto) {
	v.ServiceSpuCode = ""
	v.ServiceId = 0
	v.Quantity = 0
	poolServiceItemLineDto.Put(v)
}
