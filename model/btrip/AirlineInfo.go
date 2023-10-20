package btrip

import (
	"sync"
)

// AirlineInfo 结构体
type AirlineInfo struct {
	// 航司编码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 航司名称
	AirlineName string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// 航司简称
	AirlineSimpleName string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

var poolAirlineInfo = sync.Pool{
	New: func() any {
		return new(AirlineInfo)
	},
}

// GetAirlineInfo() 从对象池中获取AirlineInfo
func GetAirlineInfo() *AirlineInfo {
	return poolAirlineInfo.Get().(*AirlineInfo)
}

// ReleaseAirlineInfo 释放AirlineInfo
func ReleaseAirlineInfo(v *AirlineInfo) {
	v.AirlineCode = ""
	v.AirlineName = ""
	v.AirlineSimpleName = ""
	poolAirlineInfo.Put(v)
}
