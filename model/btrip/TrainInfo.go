package btrip

import (
	"sync"
)

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

var poolTrainInfo = sync.Pool{
	New: func() any {
		return new(TrainInfo)
	},
}

// GetTrainInfo() 从对象池中获取TrainInfo
func GetTrainInfo() *TrainInfo {
	return poolTrainInfo.Get().(*TrainInfo)
}

// ReleaseTrainInfo 释放TrainInfo
func ReleaseTrainInfo(v *TrainInfo) {
	v.ToStationName = ""
	v.DepTime = ""
	v.FromStationName = ""
	v.TrainNo = ""
	v.ArrTime = ""
	v.RunTime = 0
	poolTrainInfo.Put(v)
}
