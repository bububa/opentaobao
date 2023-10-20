package lstlogistics2

import (
	"sync"
)

// AddressParam 结构体
type AddressParam struct {
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 街道
	StreetName string `json:"street_name,omitempty" xml:"street_name,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
}

var poolAddressParam = sync.Pool{
	New: func() any {
		return new(AddressParam)
	},
}

// GetAddressParam() 从对象池中获取AddressParam
func GetAddressParam() *AddressParam {
	return poolAddressParam.Get().(*AddressParam)
}

// ReleaseAddressParam 释放AddressParam
func ReleaseAddressParam(v *AddressParam) {
	v.ProvinceName = ""
	v.CityName = ""
	v.DistrictName = ""
	v.StreetName = ""
	v.Address = ""
	poolAddressParam.Put(v)
}
