package alitripmerchant

import (
	"sync"
)

// AddressVo 结构体
type AddressVo struct {
	// 酒店详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 城市中文名
	CityCn string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
	// 国家中文名
	CountryCn string `json:"country_cn,omitempty" xml:"country_cn,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 地图类型
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 城市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 国内外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 国家编码
	CountryCode int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
}

var poolAddressVo = sync.Pool{
	New: func() any {
		return new(AddressVo)
	},
}

// GetAddressVo() 从对象池中获取AddressVo
func GetAddressVo() *AddressVo {
	return poolAddressVo.Get().(*AddressVo)
}

// ReleaseAddressVo 释放AddressVo
func ReleaseAddressVo(v *AddressVo) {
	v.Address = ""
	v.Latitude = ""
	v.CityCn = ""
	v.CountryCn = ""
	v.Longitude = ""
	v.PositionType = 0
	v.CityCode = 0
	v.Domestic = 0
	v.CountryCode = 0
	poolAddressVo.Put(v)
}
