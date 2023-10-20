package axindata

import (
	"sync"
)

// FlightInfoApiDto 结构体
type FlightInfoApiDto struct {
	// 行程天数第几天
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 行程日期 yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 航班号
	FlightCode string `json:"flight_code,omitempty" xml:"flight_code,omitempty"`
	// 航空公司代码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 出发城市名称
	StartCity string `json:"start_city,omitempty" xml:"start_city,omitempty"`
	// 起飞时间HHmm(时分)
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 起飞机场代码 机场三字码
	StartAirport string `json:"start_airport,omitempty" xml:"start_airport,omitempty"`
	// 抵达城市
	ArriveCity string `json:"arrive_city,omitempty" xml:"arrive_city,omitempty"`
	// 抵达时间HHmm(时分)
	ArriveTime string `json:"arrive_time,omitempty" xml:"arrive_time,omitempty"`
	// 降落机场代码
	ArriveAirport string `json:"arrive_airport,omitempty" xml:"arrive_airport,omitempty"`
	// 机型
	FlightType string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// 舱型
	CabinType string `json:"cabin_type,omitempty" xml:"cabin_type,omitempty"`
	// 是否包机 true-是false-否
	BuyOut bool `json:"buy_out,omitempty" xml:"buy_out,omitempty"`
}

var poolFlightInfoApiDto = sync.Pool{
	New: func() any {
		return new(FlightInfoApiDto)
	},
}

// GetFlightInfoApiDto() 从对象池中获取FlightInfoApiDto
func GetFlightInfoApiDto() *FlightInfoApiDto {
	return poolFlightInfoApiDto.Get().(*FlightInfoApiDto)
}

// ReleaseFlightInfoApiDto 释放FlightInfoApiDto
func ReleaseFlightInfoApiDto(v *FlightInfoApiDto) {
	v.Day = ""
	v.Date = ""
	v.FlightCode = ""
	v.AirlineCode = ""
	v.StartCity = ""
	v.StartTime = ""
	v.StartAirport = ""
	v.ArriveCity = ""
	v.ArriveTime = ""
	v.ArriveAirport = ""
	v.FlightType = ""
	v.CabinType = ""
	v.BuyOut = false
	poolFlightInfoApiDto.Put(v)
}
