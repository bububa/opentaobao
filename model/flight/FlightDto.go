package flight

// FlightDto 结构体
type FlightDto struct {
	// 允许航班日期
	AllowTravelDate []DatePairDto `json:"allow_travel_date,omitempty" xml:"allow_travel_date>date_pair_dto,omitempty"`
	// 班期
	DayWeek []string `json:"day_week,omitempty" xml:"day_week>string,omitempty"`
	// 禁止航班日期
	RestrictTravelDate []DatePairDto `json:"restrict_travel_date,omitempty" xml:"restrict_travel_date>date_pair_dto,omitempty"`
	// 可售航班日期
	AllowTravelDates []DatePairDto `json:"allow_travel_dates,omitempty" xml:"allow_travel_dates>date_pair_dto,omitempty"`
	// 不可售航班日期
	RestrictTravelDates []DatePairDto `json:"restrict_travel_dates,omitempty" xml:"restrict_travel_dates>date_pair_dto,omitempty"`
	// 班期限制
	DayWeeks []string `json:"day_weeks,omitempty" xml:"day_weeks>string,omitempty"`
	// 允许航班号，支持录入多个航班与航班范围，范围之间用“-”链接，多个用英文”,”分隔。最多支持8000字符
	AllowFlightNum string `json:"allow_flight_num,omitempty" xml:"allow_flight_num,omitempty"`
	// 允许航班起飞时间
	AllowTravelTime string `json:"allow_travel_time,omitempty" xml:"allow_travel_time,omitempty"`
	// 舱位，多个用英文”,”分隔
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 禁止航班号
	RestrictFlightNum string `json:"restrict_flight_num,omitempty" xml:"restrict_flight_num,omitempty"`
	// 第二段可售航班号
	AllowFlightNum2 string `json:"allow_flight_num2,omitempty" xml:"allow_flight_num2,omitempty"`
	// 第二段不可售航班号
	RestrictFlightNum2 string `json:"restrict_flight_num2,omitempty" xml:"restrict_flight_num2,omitempty"`
	// 行程类型标记：0，单程；1，往返
	FlightIndex int64 `json:"flight_index,omitempty" xml:"flight_index,omitempty"`
}
