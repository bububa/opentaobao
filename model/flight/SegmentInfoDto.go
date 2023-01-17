package flight

// SegmentInfoDto 结构体
type SegmentInfoDto struct {
	// 舱位信息
	CabinInfoList []CabinInfoDto `json:"cabin_info_list,omitempty" xml:"cabin_info_list>cabin_info_dto,omitempty"`
	// 抵达机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 抵达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 抵达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 起飞机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 起飞城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 起飞时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 市场航司
	MarketingAirline string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// 市场航班号
	MarketingFlightNo string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// 航程下标
	OdIndex string `json:"od_index,omitempty" xml:"od_index,omitempty"`
	// 承运航司
	OperatingAirline string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// 承运航班号
	OperatingFlightNo string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// 航段下标
	SegmentIndex string `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}
