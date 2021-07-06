package ieagency

// IeChangeItineraryVo 结构体
type IeChangeItineraryVo struct {
	// 改签航班信息
	ChangeFlights []IeChangeFlightSegmentVo `json:"change_flights,omitempty" xml:"change_flights>ie_change_flight_segment_vo,omitempty"`
	// 到达机场码
	ArrAirPortCode string `json:"arr_air_port_code,omitempty" xml:"arr_air_port_code,omitempty"`
	// 到达城市码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 到达城市名称
	ArrCityName string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// 出发机场码
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 出发城市码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 出发城市名称
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 航程序号
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
}
