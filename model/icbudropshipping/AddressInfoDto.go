package icbudropshipping

import (
	"sync"
)

// AddressInfoDto 结构体
type AddressInfoDto struct {
	// Shipping address
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// If any, please send it to us to make the freight more accurate.
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// City
	City *DivisionInfoDto `json:"city,omitempty" xml:"city,omitempty"`
	// Country
	Country *DivisionInfoDto `json:"country,omitempty" xml:"country,omitempty"`
	// province
	Province *DivisionInfoDto `json:"province,omitempty" xml:"province,omitempty"`
}

var poolAddressInfoDto = sync.Pool{
	New: func() any {
		return new(AddressInfoDto)
	},
}

// GetAddressInfoDto() 从对象池中获取AddressInfoDto
func GetAddressInfoDto() *AddressInfoDto {
	return poolAddressInfoDto.Get().(*AddressInfoDto)
}

// ReleaseAddressInfoDto 释放AddressInfoDto
func ReleaseAddressInfoDto(v *AddressInfoDto) {
	v.Address = ""
	v.Zip = ""
	v.City = nil
	v.Country = nil
	v.Province = nil
	poolAddressInfoDto.Put(v)
}
