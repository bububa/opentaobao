package btrip

import (
	"sync"
)

// JourneyRs 结构体
type JourneyRs struct {
	// 组成当前行程的航段列表
	SegmentList []FlightSegmentRs `json:"segment_list,omitempty" xml:"segment_list>flight_segment_rs,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 总时长
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 航程序号，从0开始
	SeqId int64 `json:"seq_id,omitempty" xml:"seq_id,omitempty"`
	// 总中转时间
	TransferTime int64 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
	// 是否换机场
	TransferChangeAirport bool `json:"transfer_change_airport,omitempty" xml:"transfer_change_airport,omitempty"`
}

var poolJourneyRs = sync.Pool{
	New: func() any {
		return new(JourneyRs)
	},
}

// GetJourneyRs() 从对象池中获取JourneyRs
func GetJourneyRs() *JourneyRs {
	return poolJourneyRs.Get().(*JourneyRs)
}

// ReleaseJourneyRs 释放JourneyRs
func ReleaseJourneyRs(v *JourneyRs) {
	v.SegmentList = v.SegmentList[:0]
	v.ArrCity = ""
	v.ArrTime = ""
	v.DepCity = ""
	v.DepTime = ""
	v.Duration = 0
	v.SeqId = 0
	v.TransferTime = 0
	v.TransferChangeAirport = false
	poolJourneyRs.Put(v)
}
