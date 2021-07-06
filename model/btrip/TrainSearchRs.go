package btrip

// TrainSearchRs 结构体
type TrainSearchRs struct {
	// 直达车次列表
	Trains []TrainStationVo `json:"trains,omitempty" xml:"trains>train_station_vo,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 到达站
	ArrLocation string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 出发站
	DepLocation string `json:"dep_location,omitempty" xml:"dep_location,omitempty"`
	// 是否有更多车次
	HasMoreTrain bool `json:"has_more_train,omitempty" xml:"has_more_train,omitempty"`
	// 是否展示中转引导
	ShowTransGuide bool `json:"show_trans_guide,omitempty" xml:"show_trans_guide,omitempty"`
}
