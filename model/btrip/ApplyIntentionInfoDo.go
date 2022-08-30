package btrip

// ApplyIntentionInfoDo 结构体
type ApplyIntentionInfoDo struct {
	// 出发城市名称
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// 到达城市名称
	ArrCityName string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// 出发城市三字码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 到达城市三字码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 超标的舱位，F：头等舱 C：商务舱 Y：经济舱 P：超值经济舱
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱等描述，头等舱，商务舱，经济舱，超值经济舱
	CabinClassStr string `json:"cabin_class_str,omitempty" xml:"cabin_class_str,omitempty"`
	// 折扣
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 入住城市三字码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 入住城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 入住日期
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店日期
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 到达站点名称
	ArrStation string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// 出发站点名称
	DepStation string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// 意向车次号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 意向车次类型
	TrainTypeDesc string `json:"train_type_desc,omitempty" xml:"train_type_desc,omitempty"`
	// 意向坐席名称
	SeatName string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// 申请超标的舱等 0：头等舱 1：商务舱 2：经济舱 3：超值经济舱
	CabinClass int64 `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 价格（元）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 超标类型，1:折扣 2,8,10:时间 3,9,11:折扣和时间
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否合住
	Together bool `json:"together,omitempty" xml:"together,omitempty"`
}
