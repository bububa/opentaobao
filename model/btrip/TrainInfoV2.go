package btrip

import (
	"sync"
)

// TrainInfoV2 结构体
type TrainInfoV2 struct {
	// 票信息
	TrainTicketInfos []TrainTicketInfo `json:"train_ticket_infos,omitempty" xml:"train_ticket_infos>train_ticket_info,omitempty"`
	// 出发车站名称
	FromStationName string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// 到达车站名称
	ToStationName string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 车次编号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 出发城市名
	FromCityName string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	// 到达城市名
	ToCityName string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	// 运行时长
	RunTime int64 `json:"run_time,omitempty" xml:"run_time,omitempty"`
}

var poolTrainInfoV2 = sync.Pool{
	New: func() any {
		return new(TrainInfoV2)
	},
}

// GetTrainInfoV2() 从对象池中获取TrainInfoV2
func GetTrainInfoV2() *TrainInfoV2 {
	return poolTrainInfoV2.Get().(*TrainInfoV2)
}

// ReleaseTrainInfoV2 释放TrainInfoV2
func ReleaseTrainInfoV2(v *TrainInfoV2) {
	v.TrainTicketInfos = v.TrainTicketInfos[:0]
	v.FromStationName = ""
	v.ToStationName = ""
	v.DepTime = ""
	v.TrainNo = ""
	v.ArrTime = ""
	v.FromCityName = ""
	v.ToCityName = ""
	v.RunTime = 0
	poolTrainInfoV2.Put(v)
}
