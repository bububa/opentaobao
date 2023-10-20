package tmallservice

import (
	"sync"
)

// AddressDto 结构体
type AddressDto struct {
	// 省/市/区/街道
	AddressNames []string `json:"address_names,omitempty" xml:"address_names>string,omitempty"`
	// 详细地址，街到门牌，
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
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
	v.AddressNames = v.AddressNames[:0]
	v.AddressDetail = ""
	poolAddressDto.Put(v)
}
