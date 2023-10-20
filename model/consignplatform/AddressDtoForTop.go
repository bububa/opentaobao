package consignplatform

import (
	"sync"
)

// AddressDtoForTop 结构体
type AddressDtoForTop struct {
	// 国家
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 省份
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 城市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 街道
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
}

var poolAddressDtoForTop = sync.Pool{
	New: func() any {
		return new(AddressDtoForTop)
	},
}

// GetAddressDtoForTop() 从对象池中获取AddressDtoForTop
func GetAddressDtoForTop() *AddressDtoForTop {
	return poolAddressDtoForTop.Get().(*AddressDtoForTop)
}

// ReleaseAddressDtoForTop 释放AddressDtoForTop
func ReleaseAddressDtoForTop(v *AddressDtoForTop) {
	v.CountryName = ""
	v.ProvName = ""
	v.CityName = ""
	v.AreaName = ""
	v.TownName = ""
	v.AddressDetail = ""
	poolAddressDtoForTop.Put(v)
}
