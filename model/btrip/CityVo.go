package btrip

import (
	"sync"
)

// CityVo 结构体
type CityVo struct {
	// 三字码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 城市名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 所属省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 邻近机场城市，只在邻近机场推荐有值
	TravelName string `json:"travel_name,omitempty" xml:"travel_name,omitempty"`
	// 与搜索城市距离，单位千米，只在邻近机场推荐有值
	Distance int64 `json:"distance,omitempty" xml:"distance,omitempty"`
}

var poolCityVo = sync.Pool{
	New: func() any {
		return new(CityVo)
	},
}

// GetCityVo() 从对象池中获取CityVo
func GetCityVo() *CityVo {
	return poolCityVo.Get().(*CityVo)
}

// ReleaseCityVo 释放CityVo
func ReleaseCityVo(v *CityVo) {
	v.Code = ""
	v.Name = ""
	v.City = ""
	v.Province = ""
	v.TravelName = ""
	v.Distance = 0
	poolCityVo.Put(v)
}
