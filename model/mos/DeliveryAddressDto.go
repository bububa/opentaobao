package mos

import (
	"sync"
)

// DeliveryAddressDto 结构体
type DeliveryAddressDto struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细信息
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 编码
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
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
	v.Province = ""
	v.City = ""
	v.District = ""
	v.Town = ""
	v.DetailAddress = ""
	v.ZipCode = ""
	v.DivisionId = 0
	poolDeliveryAddressDto.Put(v)
}
