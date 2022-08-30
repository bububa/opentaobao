package ieagency

// IeBookFlightVo 结构体
type IeBookFlightVo struct {
	// 到达机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 出发机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 出发日期
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 航班方向(Outbound:去程,Inbound:回程,行程类型,MultiCity:多程)
	DirectionType string `json:"direction_type,omitempty" xml:"direction_type,omitempty"`
	// 机型
	EquipType string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// 航班舱位
	FlightCabin string `json:"flight_cabin,omitempty" xml:"flight_cabin,omitempty"`
	// 市场方航班号
	FlightNumber string `json:"flight_number,omitempty" xml:"flight_number,omitempty"`
	// 市场
	MarketingAirline string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// 共享航空公司
	OperatingAirLine string `json:"operating_air_line,omitempty" xml:"operating_air_line,omitempty"`
	// 共享航班号
	OperatingFlightNumber string `json:"operating_flight_number,omitempty" xml:"operating_flight_number,omitempty"`
	// 第几段
	SegmentRph int64 `json:"segment_rph,omitempty" xml:"segment_rph,omitempty"`
	// 是否共享航班
	CodeShare bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
}
