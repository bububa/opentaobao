package cainiaolocker

import (
	"sync"
)

// AddressDto 结构体
type AddressDto struct {
	// 城市，长度小于20
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 详细地址，长度小于256
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 区，长度小于20
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 省，长度小于20
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道，长度小于30
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}

var poolAddressDto = sync.Pool{
	New: func() any {
		return new(AddressDto)
	},
}

// GetAddressDto() 从对象池中获取AddressDto
func GetAddressDto() *AddressDto {
	return poolAddressDto.Get().(*AddressDto)
}

// ReleaseAddressDto 释放AddressDto
func ReleaseAddressDto(v *AddressDto) {
	v.City = ""
	v.Detail = ""
	v.District = ""
	v.Province = ""
	v.Town = ""
	poolAddressDto.Put(v)
}
