package eleenterpriseordernew

import (
	"sync"
)

// DeliverymanInfoDto 结构体
type DeliverymanInfoDto struct {
	// 配送员姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 配送员电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolDeliverymanInfoDto = sync.Pool{
	New: func() any {
		return new(DeliverymanInfoDto)
	},
}

// GetDeliverymanInfoDto() 从对象池中获取DeliverymanInfoDto
func GetDeliverymanInfoDto() *DeliverymanInfoDto {
	return poolDeliverymanInfoDto.Get().(*DeliverymanInfoDto)
}

// ReleaseDeliverymanInfoDto 释放DeliverymanInfoDto
func ReleaseDeliverymanInfoDto(v *DeliverymanInfoDto) {
	v.Name = ""
	v.Phone = ""
	poolDeliverymanInfoDto.Put(v)
}
