package examination

import (
	"sync"
)

// AddAddress 结构体
type AddAddress struct {
	// 用户的上门地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 上门地址纬度（高德系）
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 上门地址经度（高德系）
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 区域编码（高德标准）
	DistrictCode string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 区域名称（高德标准）
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 城市编码（高德标准）
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称（高德标准）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省份编码（高德标准）
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 省份名称（高德标准）
	Province string `json:"province,omitempty" xml:"province,omitempty"`
}

var poolAddAddress = sync.Pool{
	New: func() any {
		return new(AddAddress)
	},
}

// GetAddAddress() 从对象池中获取AddAddress
func GetAddAddress() *AddAddress {
	return poolAddAddress.Get().(*AddAddress)
}

// ReleaseAddAddress 释放AddAddress
func ReleaseAddAddress(v *AddAddress) {
	v.Address = ""
	v.Latitude = ""
	v.Longitude = ""
	v.DistrictCode = ""
	v.District = ""
	v.CityCode = ""
	v.City = ""
	v.ProvinceCode = ""
	v.Province = ""
	poolAddAddress.Put(v)
}
