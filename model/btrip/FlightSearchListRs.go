package btrip

import (
	"sync"
)

// FlightSearchListRs 结构体
type FlightSearchListRs struct {
	// 航班列表
	FlightList []FlightInfoDto `json:"flight_list,omitempty" xml:"flight_list>flight_info_dto,omitempty"`
	// 是否可更换PNR出票
	IsReplacePnr bool `json:"is_replace_pnr,omitempty" xml:"is_replace_pnr,omitempty"`
}

var poolFlightSearchListRs = sync.Pool{
	New: func() any {
		return new(FlightSearchListRs)
	},
}

// GetFlightSearchListRs() 从对象池中获取FlightSearchListRs
func GetFlightSearchListRs() *FlightSearchListRs {
	return poolFlightSearchListRs.Get().(*FlightSearchListRs)
}

// ReleaseFlightSearchListRs 释放FlightSearchListRs
func ReleaseFlightSearchListRs(v *FlightSearchListRs) {
	v.FlightList = v.FlightList[:0]
	v.IsReplacePnr = false
	poolFlightSearchListRs.Put(v)
}
