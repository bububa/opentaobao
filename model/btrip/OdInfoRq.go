package btrip

// OdInfoRq 结构体
type OdInfoRq struct {
	// 到达机场三字码
	ArrAirportCode []string `json:"arr_airport_code,omitempty" xml:"arr_airport_code>string,omitempty"`
	// 出发机场三字码
	DepAirportCode []string `json:"dep_airport_code,omitempty" xml:"dep_airport_code>string,omitempty"`
	// 到达城市三字码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 出发城市三字码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 起飞日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 最早起飞时间
	EarliestDepTime string `json:"earliest_dep_time,omitempty" xml:"earliest_dep_time,omitempty"`
	// 最晚起飞时间
	LatestDepTime string `json:"latest_dep_time,omitempty" xml:"latest_dep_time,omitempty"`
}
