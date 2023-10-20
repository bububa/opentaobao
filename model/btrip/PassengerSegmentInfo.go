package btrip

import (
	"sync"
)

// PassengerSegmentInfo 结构体
type PassengerSegmentInfo struct {
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 用户编号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolPassengerSegmentInfo = sync.Pool{
	New: func() any {
		return new(PassengerSegmentInfo)
	},
}

// GetPassengerSegmentInfo() 从对象池中获取PassengerSegmentInfo
func GetPassengerSegmentInfo() *PassengerSegmentInfo {
	return poolPassengerSegmentInfo.Get().(*PassengerSegmentInfo)
}

// ReleasePassengerSegmentInfo 释放PassengerSegmentInfo
func ReleasePassengerSegmentInfo(v *PassengerSegmentInfo) {
	v.FlightNo = ""
	v.PassengerName = ""
	v.UserId = ""
	poolPassengerSegmentInfo.Put(v)
}
