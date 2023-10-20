package idle

import (
	"sync"
)

// AddressDto 结构体
type AddressDto struct {
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
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
	v.Area = ""
	v.City = ""
	v.Prov = ""
	poolAddressDto.Put(v)
}
