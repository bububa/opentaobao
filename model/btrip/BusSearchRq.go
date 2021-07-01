package btrip

// BusSearchRq 结构体
type BusSearchRq struct {
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 目的地名称过虑条件。多个用都还分隔，如：临湘,房县,枣阳
	ArriveCityName string `json:"arrive_city_name,omitempty" xml:"arrive_city_name,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 出发车次名称过虑条件,多个用逗号分隔，如：省汽车客运站,天河客运站,越秀南客运站
	FromStationName string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// 针对发车时段过滤新增字段：1:凌晨 2:上午 3:下午 4:晚上,多个用逗号分隔
	PeriodTime string `json:"period_time,omitempty" xml:"period_time,omitempty"`
	// 汽车票场景 普通 common 机场巴士 airport
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 是否显示不可售车次
	ShowNoSell int64 `json:"show_no_sell,omitempty" xml:"show_no_sell,omitempty"`
	// 标准出发城市编码
	StandardFromAreaCode string `json:"standard_from_area_code,omitempty" xml:"standard_from_area_code,omitempty"`
	// 标准出发站编码
	StandardFromStationId int64 `json:"standard_from_station_id,omitempty" xml:"standard_from_station_id,omitempty"`
	// 出发车站名称
	StandardFromStationName string `json:"standard_from_station_name,omitempty" xml:"standard_from_station_name,omitempty"`
	// 标准到达城市编码
	StandardToAreaCode string `json:"standard_to_area_code,omitempty" xml:"standard_to_area_code,omitempty"`
	// 针对到达站过滤新增字段,多个用逗号分隔
	ToStationName string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
}
