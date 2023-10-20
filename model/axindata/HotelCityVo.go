package axindata

import (
	"sync"
)

// HotelCityVo 结构体
type HotelCityVo struct {
	// 城市名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 城市编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolHotelCityVo = sync.Pool{
	New: func() any {
		return new(HotelCityVo)
	},
}

// GetHotelCityVo() 从对象池中获取HotelCityVo
func GetHotelCityVo() *HotelCityVo {
	return poolHotelCityVo.Get().(*HotelCityVo)
}

// ReleaseHotelCityVo 释放HotelCityVo
func ReleaseHotelCityVo(v *HotelCityVo) {
	v.Name = ""
	v.Code = 0
	poolHotelCityVo.Put(v)
}
