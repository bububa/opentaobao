package tmallcar

import (
	"sync"
)

// DeliveryAddressDto 结构体
type DeliveryAddressDto struct {
	// 收货人详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}

var poolDeliveryAddressDto = sync.Pool{
	New: func() any {
		return new(DeliveryAddressDto)
	},
}

// GetDeliveryAddressDto() 从对象池中获取DeliveryAddressDto
func GetDeliveryAddressDto() *DeliveryAddressDto {
	return poolDeliveryAddressDto.Get().(*DeliveryAddressDto)
}

// ReleaseDeliveryAddressDto 释放DeliveryAddressDto
func ReleaseDeliveryAddressDto(v *DeliveryAddressDto) {
	v.DetailAddress = ""
	poolDeliveryAddressDto.Put(v)
}
