package btrip

import (
	"sync"
)

// BtripFlightSearchListRs 结构体
type BtripFlightSearchListRs struct {
	// 航班列表
	FlightInfoList []FlightInfoDto `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info_dto,omitempty"`
	// 是否可更换PNR出票
	IsReplacePnr bool `json:"is_replace_pnr,omitempty" xml:"is_replace_pnr,omitempty"`
}

var poolBtripFlightSearchListRs = sync.Pool{
	New: func() any {
		return new(BtripFlightSearchListRs)
	},
}

// GetBtripFlightSearchListRs() 从对象池中获取BtripFlightSearchListRs
func GetBtripFlightSearchListRs() *BtripFlightSearchListRs {
	return poolBtripFlightSearchListRs.Get().(*BtripFlightSearchListRs)
}

// ReleaseBtripFlightSearchListRs 释放BtripFlightSearchListRs
func ReleaseBtripFlightSearchListRs(v *BtripFlightSearchListRs) {
	v.FlightInfoList = v.FlightInfoList[:0]
	v.IsReplacePnr = false
	poolBtripFlightSearchListRs.Put(v)
}
