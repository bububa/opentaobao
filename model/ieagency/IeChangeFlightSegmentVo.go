package ieagency

// IeChangeFlightSegmentVo 结构体
type IeChangeFlightSegmentVo struct {
	// 到达机场码
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 到达城市码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 到达航站楼
	ArrTerminal string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// 到达日期
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 舱位等级
	CabinClassCode string `json:"cabin_class_code,omitempty" xml:"cabin_class_code,omitempty"`
	// 舱位代码
	CabinCode string `json:"cabin_code,omitempty" xml:"cabin_code,omitempty"`
	// 是否共享
	CodeShare bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// 出发机场码
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 出发城市码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发航站楼
	DepTerminal string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 市场承运航班号
	MarketingFlightNumber string `json:"marketing_flight_number,omitempty" xml:"marketing_flight_number,omitempty"`
	// 实际承运航班号
	OperatingFlightNumber string `json:"operating_flight_number,omitempty" xml:"operating_flight_number,omitempty"`
	// 航段序号
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}
