package btrip

import (
	"sync"
)

// AirportInfo 结构体
type AirportInfo struct {
	// 机场编码
	AirportCode string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	// 机场名称
	AirportName string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// 航站楼
	Terminal string `json:"terminal,omitempty" xml:"terminal,omitempty"`
	// 城市三字码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

var poolAirportInfo = sync.Pool{
	New: func() any {
		return new(AirportInfo)
	},
}

// GetAirportInfo() 从对象池中获取AirportInfo
func GetAirportInfo() *AirportInfo {
	return poolAirportInfo.Get().(*AirportInfo)
}

// ReleaseAirportInfo 释放AirportInfo
func ReleaseAirportInfo(v *AirportInfo) {
	v.AirportCode = ""
	v.AirportName = ""
	v.Terminal = ""
	v.CityCode = ""
	v.CityName = ""
	poolAirportInfo.Put(v)
}
