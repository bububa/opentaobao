package maitix

import (
	"sync"
)

// AddressDto 结构体
type AddressDto struct {
	// 区域国标ID-可以不填
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 城市国标ID-必填
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 省国标ID-可以不填
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 国家国标ID-可以不填
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
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
	v.AreaId = 0
	v.CityId = 0
	v.ProvinceId = 0
	v.CountryId = 0
	poolAddressDto.Put(v)
}
