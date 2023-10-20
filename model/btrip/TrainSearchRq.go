package btrip

import (
	"sync"
)

// TrainSearchRq 结构体
type TrainSearchRq struct {
	// 目的地区域code
	ArrAreaCode string `json:"arr_area_code,omitempty" xml:"arr_area_code,omitempty"`
	// 目的地地区
	ArrAreaName string `json:"arr_area_name,omitempty" xml:"arr_area_name,omitempty"`
	// 目的地
	ArrLocation string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	// 目的地code，站点三字码
	ArrLocationCode string `json:"arr_location_code,omitempty" xml:"arr_location_code,omitempty"`
	// 出发地区
	DepAreaCode string `json:"dep_area_code,omitempty" xml:"dep_area_code,omitempty"`
	// 出发地区
	DepAreaName string `json:"dep_area_name,omitempty" xml:"dep_area_name,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 出发地
	DepLocation string `json:"dep_location,omitempty" xml:"dep_location,omitempty"`
	// 出发地code，站点三字码
	DepLocationCode string `json:"dep_location_code,omitempty" xml:"dep_location_code,omitempty"`
	// 企业corpId
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 0:普通(默认) 1：学生
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 排序规则，排序规则:0:最早出发,1:最晚出发,2:耗时最短
	SortType int64 `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
}

var poolTrainSearchRq = sync.Pool{
	New: func() any {
		return new(TrainSearchRq)
	},
}

// GetTrainSearchRq() 从对象池中获取TrainSearchRq
func GetTrainSearchRq() *TrainSearchRq {
	return poolTrainSearchRq.Get().(*TrainSearchRq)
}

// ReleaseTrainSearchRq 释放TrainSearchRq
func ReleaseTrainSearchRq(v *TrainSearchRq) {
	v.ArrAreaCode = ""
	v.ArrAreaName = ""
	v.ArrLocation = ""
	v.ArrLocationCode = ""
	v.DepAreaCode = ""
	v.DepAreaName = ""
	v.DepDate = ""
	v.DepLocation = ""
	v.DepLocationCode = ""
	v.CorpId = ""
	v.PassengerType = 0
	v.SortType = 0
	poolTrainSearchRq.Put(v)
}
