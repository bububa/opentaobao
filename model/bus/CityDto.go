package bus

import (
	"sync"
)

// CityDto 结构体
type CityDto struct {
	// 城市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 拼音
	CityFullPy string `json:"city_full_py,omitempty" xml:"city_full_py,omitempty"`
	// 城市名字
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 简拼
	CityShortPy string `json:"city_short_py,omitempty" xml:"city_short_py,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 预售期
	PreDay int64 `json:"pre_day,omitempty" xml:"pre_day,omitempty"`
}

var poolCityDto = sync.Pool{
	New: func() any {
		return new(CityDto)
	},
}

// GetCityDto() 从对象池中获取CityDto
func GetCityDto() *CityDto {
	return poolCityDto.Get().(*CityDto)
}

// ReleaseCityDto 释放CityDto
func ReleaseCityDto(v *CityDto) {
	v.CityCode = ""
	v.CityFullPy = ""
	v.CityName = ""
	v.CityShortPy = ""
	v.ProvinceName = ""
	v.PreDay = 0
	poolCityDto.Put(v)
}
