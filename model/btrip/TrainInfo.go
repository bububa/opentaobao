package btrip

// TrainInfo 结构体
type TrainInfo struct {
	// 到达车站名称
	ToStationName string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 出发车站名称
	FromStationName string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// 车次编号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 运行时长
	RunTime int64 `json:"run_time,omitempty" xml:"run_time,omitempty"`
}
