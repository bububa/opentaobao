package alitripmerchant

import (
	"sync"
)

// AddressListSearchDto 结构体
type AddressListSearchDto struct {
	// 城市列表
	CityList []AddressLetterDto `json:"city_list,omitempty" xml:"city_list>address_letter_dto,omitempty"`
	// 热门城市
	HotCityList []AddressSearchDto `json:"hot_city_list,omitempty" xml:"hot_city_list>address_search_dto,omitempty"`
}

var poolAddressListSearchDto = sync.Pool{
	New: func() any {
		return new(AddressListSearchDto)
	},
}

// GetAddressListSearchDto() 从对象池中获取AddressListSearchDto
func GetAddressListSearchDto() *AddressListSearchDto {
	return poolAddressListSearchDto.Get().(*AddressListSearchDto)
}

// ReleaseAddressListSearchDto 释放AddressListSearchDto
func ReleaseAddressListSearchDto(v *AddressListSearchDto) {
	v.CityList = v.CityList[:0]
	v.HotCityList = v.HotCityList[:0]
	poolAddressListSearchDto.Put(v)
}
