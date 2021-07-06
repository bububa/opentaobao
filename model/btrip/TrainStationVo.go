package btrip

// TrainStationVo 结构体
type TrainStationVo struct {
	// 坐席列表
	SeatTypes []SeatVo `json:"seat_types,omitempty" xml:"seat_types>seat_vo,omitempty"`
	// 次日
	ArrDayTag string `json:"arr_day_tag,omitempty" xml:"arr_day_tag,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 完整到达时间,
	ArrTimeFull string `json:"arr_time_full,omitempty" xml:"arr_time_full,omitempty"`
	// 到达车站站名,
	ArriveStation string `json:"arrive_station,omitempty" xml:"arrive_station,omitempty"`
	// 到达车站三字码,
	ArriveStationCode string `json:"arrive_station_code,omitempty" xml:"arrive_station_code,omitempty"`
	// 耗时,
	CostTime string `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// 出发时间,
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 完整出发时间
	DepTimeFull string `json:"dep_time_full,omitempty" xml:"dep_time_full,omitempty"`
	// 出发车站站名,
	DepartStation string `json:"depart_station,omitempty" xml:"depart_station,omitempty"`
	// 出发车站三字码,
	DepartStationCode string `json:"depart_station_code,omitempty" xml:"depart_station_code,omitempty"`
	// 预售期
	PreSellDateStr string `json:"pre_sell_date_str,omitempty" xml:"pre_sell_date_str,omitempty"`
	// 列车车次
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 耗时,分钟
	CostTimeInt int64 `json:"cost_time_int,omitempty" xml:"cost_time_int,omitempty"`
	// 是否终点站
	EndStation int64 `json:"end_station,omitempty" xml:"end_station,omitempty"`
	// 列表页用户显示的坐席名称对应的价格。
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 是否始发站
	StartStation int64 `json:"start_station,omitempty" xml:"start_station,omitempty"`
	// 列车类型的编号。1=普快, 2=新空普快, 3=普客, 4=快速, 5=新空普客, 6=城际高速, 7=动车组, 8=高速动车, 9=新空快速, 10=新空特快, 11=特快, 12=新空直达
	TrainType int64 `json:"train_type,omitempty" xml:"train_type,omitempty"`
	// 是否有更多车次
	HasMoreTrain bool `json:"has_more_train,omitempty" xml:"has_more_train,omitempty"`
	// 是否有票, 0:无票，1:有票
	HasStock bool `json:"has_stock,omitempty" xml:"has_stock,omitempty"`
	// 是否是复兴号
	RevivalTrain bool `json:"revival_train,omitempty" xml:"revival_train,omitempty"`
	// 是否支持身份
	ShowIdIcon bool `json:"show_id_icon,omitempty" xml:"show_id_icon,omitempty"`
	// 是否是静音车厢
	SilenceCompartment bool `json:"silence_compartment,omitempty" xml:"silence_compartment,omitempty"`
}
