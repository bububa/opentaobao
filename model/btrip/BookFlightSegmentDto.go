package btrip

// BookFlightSegmentDto 结构体
type BookFlightSegmentDto struct {
	// 到达机场
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 到达城市
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 出发机场
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 出发城市
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 航段序号
	SegmentNumber string `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	// 航段核对参数
	SegSecretParams string `json:"seg_secret_params,omitempty" xml:"seg_secret_params,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 行程单类型
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
}
