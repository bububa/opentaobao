package btrip

import (
	"sync"
)

// FlightInfo 结构体
type FlightInfo struct {
	// 到达城市
	ArrCityName string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// 起飞时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 出发机场三字码
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 舱等
	CabinLevel string `json:"cabin_level,omitempty" xml:"cabin_level,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 到达城市三字码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 出发城市
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// 航司编码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 出发机场名称
	DepAirportName string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	// 到达机场名称
	ArrAirportName string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 达到机场三字码
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 航司名称
	AirlineName string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// 出发城市三字码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 飞行里程
	FlightMile int64 `json:"flight_mile,omitempty" xml:"flight_mile,omitempty"`
}

var poolFlightInfo = sync.Pool{
	New: func() any {
		return new(FlightInfo)
	},
}

// GetFlightInfo() 从对象池中获取FlightInfo
func GetFlightInfo() *FlightInfo {
	return poolFlightInfo.Get().(*FlightInfo)
}

// ReleaseFlightInfo 释放FlightInfo
func ReleaseFlightInfo(v *FlightInfo) {
	v.ArrCityName = ""
	v.DepTime = ""
	v.Cabin = ""
	v.DepAirportCode = ""
	v.CabinLevel = ""
	v.FlightNo = ""
	v.ArrCityCode = ""
	v.DepCityName = ""
	v.AirlineCode = ""
	v.DepAirportName = ""
	v.ArrAirportName = ""
	v.ArrTime = ""
	v.ArrAirportCode = ""
	v.AirlineName = ""
	v.DepCityCode = ""
	v.FlightMile = 0
	poolFlightInfo.Put(v)
}
