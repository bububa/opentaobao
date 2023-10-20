package alitripmerchant

import (
	"sync"
)

// AddressSearchVo 结构体
type AddressSearchVo struct {
	// 酒店详情
	HotelList []HotelAddressDetail `json:"hotel_list,omitempty" xml:"hotel_list>hotel_address_detail,omitempty"`
	// 城市列表
	CityList []CityAddressDetail `json:"city_list,omitempty" xml:"city_list>city_address_detail,omitempty"`
}

var poolAddressSearchVo = sync.Pool{
	New: func() any {
		return new(AddressSearchVo)
	},
}

// GetAddressSearchVo() 从对象池中获取AddressSearchVo
func GetAddressSearchVo() *AddressSearchVo {
	return poolAddressSearchVo.Get().(*AddressSearchVo)
}

// ReleaseAddressSearchVo 释放AddressSearchVo
func ReleaseAddressSearchVo(v *AddressSearchVo) {
	v.HotelList = v.HotelList[:0]
	v.CityList = v.CityList[:0]
	poolAddressSearchVo.Put(v)
}
