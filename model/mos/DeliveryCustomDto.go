package mos

import (
	"sync"
)

// DeliveryCustomDto 结构体
type DeliveryCustomDto struct {
	// 名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 头像
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 电话
	Telphone string `json:"telphone,omitempty" xml:"telphone,omitempty"`
	// 详细信息
	AddressInfo *DeliveryAddressDto `json:"address_info,omitempty" xml:"address_info,omitempty"`
}

var poolDeliveryCustomDto = sync.Pool{
	New: func() any {
		return new(DeliveryCustomDto)
	},
}

// GetDeliveryCustomDto() 从对象池中获取DeliveryCustomDto
func GetDeliveryCustomDto() *DeliveryCustomDto {
	return poolDeliveryCustomDto.Get().(*DeliveryCustomDto)
}

// ReleaseDeliveryCustomDto 释放DeliveryCustomDto
func ReleaseDeliveryCustomDto(v *DeliveryCustomDto) {
	v.Name = ""
	v.AvatarUrl = ""
	v.MobilePhone = ""
	v.Telphone = ""
	v.AddressInfo = nil
	poolDeliveryCustomDto.Put(v)
}
