package flight

// FlightInfoDto 结构体
type FlightInfoDto struct {
	// 航段信息
	SegmentInfos []SegmentInfoDto `json:"segment_infos,omitempty" xml:"segment_infos>segment_info_dto,omitempty"`
	// 航司
	Airway string `json:"airway,omitempty" xml:"airway,omitempty"`
	// 抵达机场
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 起飞机场
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 第一段起飞时间
	DepDateTime string `json:"dep_date_time,omitempty" xml:"dep_date_time,omitempty"`
}
