package tblogistics

import (
	"sync"
)

// DeliveryTopDto 结构体
type DeliveryTopDto struct {
	// 配送员电话，支持手机、座机、400电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 配送员姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolDeliveryTopDto = sync.Pool{
	New: func() any {
		return new(DeliveryTopDto)
	},
}

// GetDeliveryTopDto() 从对象池中获取DeliveryTopDto
func GetDeliveryTopDto() *DeliveryTopDto {
	return poolDeliveryTopDto.Get().(*DeliveryTopDto)
}

// ReleaseDeliveryTopDto 释放DeliveryTopDto
func ReleaseDeliveryTopDto(v *DeliveryTopDto) {
	v.Phone = ""
	v.Name = ""
	poolDeliveryTopDto.Put(v)
}
