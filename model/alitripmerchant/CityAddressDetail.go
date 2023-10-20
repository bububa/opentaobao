package alitripmerchant

import (
	"sync"
)

// CityAddressDetail 结构体
type CityAddressDetail struct {
	// 国家(汉字)
	CountryCn string `json:"country_cn,omitempty" xml:"country_cn,omitempty"`
	// 城市(汉字)
	CityCn string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
	// 省份(汉字)
	ProvinceCn string `json:"province_cn,omitempty" xml:"province_cn,omitempty"`
	// 城市拼音首字母
	CityPyHead string `json:"city_py_head,omitempty" xml:"city_py_head,omitempty"`
	// 城市图片URL
	CityUrl string `json:"city_url,omitempty" xml:"city_url,omitempty"`
	// 0国内1国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 国家编码
	CountryCode int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 城市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 省份代码
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
}

var poolCityAddressDetail = sync.Pool{
	New: func() any {
		return new(CityAddressDetail)
	},
}

// GetCityAddressDetail() 从对象池中获取CityAddressDetail
func GetCityAddressDetail() *CityAddressDetail {
	return poolCityAddressDetail.Get().(*CityAddressDetail)
}

// ReleaseCityAddressDetail 释放CityAddressDetail
func ReleaseCityAddressDetail(v *CityAddressDetail) {
	v.CountryCn = ""
	v.CityCn = ""
	v.ProvinceCn = ""
	v.CityPyHead = ""
	v.CityUrl = ""
	v.Domestic = 0
	v.CountryCode = 0
	v.CityCode = 0
	v.ProvinceCode = 0
	poolCityAddressDetail.Put(v)
}
