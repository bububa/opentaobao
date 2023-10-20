package btrip

import (
	"sync"
)

// ModifyFlightInfo 结构体
type ModifyFlightInfo struct {
	// 改签乘客信息列表
	PassengerInfoList []PassengerCabinInfo `json:"passenger_info_list,omitempty" xml:"passenger_info_list>passenger_cabin_info,omitempty"`
	// 到达城市编码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 出发城市编码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 改签航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

var poolModifyFlightInfo = sync.Pool{
	New: func() any {
		return new(ModifyFlightInfo)
	},
}

// GetModifyFlightInfo() 从对象池中获取ModifyFlightInfo
func GetModifyFlightInfo() *ModifyFlightInfo {
	return poolModifyFlightInfo.Get().(*ModifyFlightInfo)
}

// ReleaseModifyFlightInfo 释放ModifyFlightInfo
func ReleaseModifyFlightInfo(v *ModifyFlightInfo) {
	v.PassengerInfoList = v.PassengerInfoList[:0]
	v.ArrCity = ""
	v.Cabin = ""
	v.DepCity = ""
	v.DepDate = ""
	v.FlightNo = ""
	poolModifyFlightInfo.Put(v)
}
