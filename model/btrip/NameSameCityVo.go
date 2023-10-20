package btrip

import (
	"sync"
)

// NameSameCityVo 结构体
type NameSameCityVo struct {
	// 来自城市
	FromCityList []CityVo `json:"from_city_list,omitempty" xml:"from_city_list>city_vo,omitempty"`
	// 到达城市
	ToCityList []CityVo `json:"to_city_list,omitempty" xml:"to_city_list>city_vo,omitempty"`
}

var poolNameSameCityVo = sync.Pool{
	New: func() any {
		return new(NameSameCityVo)
	},
}

// GetNameSameCityVo() 从对象池中获取NameSameCityVo
func GetNameSameCityVo() *NameSameCityVo {
	return poolNameSameCityVo.Get().(*NameSameCityVo)
}

// ReleaseNameSameCityVo 释放NameSameCityVo
func ReleaseNameSameCityVo(v *NameSameCityVo) {
	v.FromCityList = v.FromCityList[:0]
	v.ToCityList = v.ToCityList[:0]
	poolNameSameCityVo.Put(v)
}
