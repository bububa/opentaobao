package alitripmerchant

import (
	"sync"
)

// AddressSearchDto 结构体
type AddressSearchDto struct {
	// 城市图片url
	CityUrl string `json:"city_url,omitempty" xml:"city_url,omitempty"`
	// 城市拼音首字母
	CityPyHead string `json:"city_py_head,omitempty" xml:"city_py_head,omitempty"`
	// 城市
	CityCn string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
	// 国家
	CountryCn string `json:"country_cn,omitempty" xml:"country_cn,omitempty"`
	// 0国内1国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 城市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 国家编码
	CountryCode int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
}

var poolAddressSearchDto = sync.Pool{
	New: func() any {
		return new(AddressSearchDto)
	},
}

// GetAddressSearchDto() 从对象池中获取AddressSearchDto
func GetAddressSearchDto() *AddressSearchDto {
	return poolAddressSearchDto.Get().(*AddressSearchDto)
}

// ReleaseAddressSearchDto 释放AddressSearchDto
func ReleaseAddressSearchDto(v *AddressSearchDto) {
	v.CityUrl = ""
	v.CityPyHead = ""
	v.CityCn = ""
	v.CountryCn = ""
	v.Domestic = 0
	v.CityCode = 0
	v.CountryCode = 0
	poolAddressSearchDto.Put(v)
}
