package jipiao

import (
	"sync"
)

// TripFlightInfo 结构体
type TripFlightInfo struct {
	// 乘机人信息列表
	Passengers []TripFlightPassenger `json:"passengers,omitempty" xml:"passengers>trip_flight_passenger,omitempty"`
	// 航班航空公司二字码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 航班实际承运航班号
	Carrier string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 航班机型
	FlightType string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// 航班出发城市三字码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 航班到达城市三字码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 航班出发机场三字码
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 航班到达机场三字码
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 航班起飞时间，格式yyyy-mm-dd hh:mm:ss
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 航班到达时间，格式yyyy-mm-dd hh:mm:ss
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 特殊说明
	SpecialRule string `json:"special_rule,omitempty" xml:"special_rule,omitempty"`
	// 淘宝机票平台航班id
	FlightId int64 `json:"flight_id,omitempty" xml:"flight_id,omitempty"`
	// 航段类型：0，去程；1，回程
	SegmentType int64 `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	// 去程或回程第几段航班，0，第1段；1，第2段航班
	SegmentNumber int64 `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	// 当前航段票面价格，单位：分
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

var poolTripFlightInfo = sync.Pool{
	New: func() any {
		return new(TripFlightInfo)
	},
}

// GetTripFlightInfo() 从对象池中获取TripFlightInfo
func GetTripFlightInfo() *TripFlightInfo {
	return poolTripFlightInfo.Get().(*TripFlightInfo)
}

// ReleaseTripFlightInfo 释放TripFlightInfo
func ReleaseTripFlightInfo(v *TripFlightInfo) {
	v.Passengers = v.Passengers[:0]
	v.AirlineCode = ""
	v.FlightNo = ""
	v.Carrier = ""
	v.FlightType = ""
	v.DepCityCode = ""
	v.ArrCityCode = ""
	v.DepAirportCode = ""
	v.ArrAirportCode = ""
	v.DepTime = ""
	v.ArrTime = ""
	v.Extra = ""
	v.SpecialRule = ""
	v.FlightId = 0
	v.SegmentType = 0
	v.SegmentNumber = 0
	v.TicketPrice = 0
	poolTripFlightInfo.Put(v)
}
