package btrip

// OrderFlightInfo 结构体
type OrderFlightInfo struct {
	// 航司编码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 到达机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 到达机场编码
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 到达城市名称
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 到达城市编码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 行李额
	Baggage string `json:"baggage,omitempty" xml:"baggage,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 承运航班
	Carrier string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 出发机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 出发机场编码
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发城市编码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 餐食
	Meal string `json:"meal,omitempty" xml:"meal,omitempty"`
	// 经停到达时间
	StopArrTime string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// 经停城市
	StopCity string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// 经停出发时间
	StopDepTime string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// 改签前航班号
	LastFlightNo string `json:"last_flight_no,omitempty" xml:"last_flight_no,omitempty"`
	// 改签前舱位
	LastCabin string `json:"last_cabin,omitempty" xml:"last_cabin,omitempty"`
	// 改签后的退改签文案
	TuigaiqianInfo string `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	// 出发航站楼
	DepTerminal string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// 到达航站楼
	ArrTerminal string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// 航司名称
	AirlineName string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// 航司简称
	AirlineSimpleName string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
	// 到达机场名称
	ArrAirportCodeName string `json:"arr_airport_code_name,omitempty" xml:"arr_airport_code_name,omitempty"`
	// 出发机场名称
	DepAirportCodeName string `json:"dep_airport_code_name,omitempty" xml:"dep_airport_code_name,omitempty"`
	// 机建费用
	BuildPrice int64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// 燃油费
	OilPrice int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// 航段类型
	SegmentType int64 `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	// 票费用
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}
