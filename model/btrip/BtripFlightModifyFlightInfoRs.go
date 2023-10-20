package btrip

import (
	"sync"
)

// BtripFlightModifyFlightInfoRs 结构体
type BtripFlightModifyFlightInfoRs struct {
	// 改签航班信息列表
	FlightInfoList []FlightInfoDto `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info_dto,omitempty"`
}

var poolBtripFlightModifyFlightInfoRs = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyFlightInfoRs)
	},
}

// GetBtripFlightModifyFlightInfoRs() 从对象池中获取BtripFlightModifyFlightInfoRs
func GetBtripFlightModifyFlightInfoRs() *BtripFlightModifyFlightInfoRs {
	return poolBtripFlightModifyFlightInfoRs.Get().(*BtripFlightModifyFlightInfoRs)
}

// ReleaseBtripFlightModifyFlightInfoRs 释放BtripFlightModifyFlightInfoRs
func ReleaseBtripFlightModifyFlightInfoRs(v *BtripFlightModifyFlightInfoRs) {
	v.FlightInfoList = v.FlightInfoList[:0]
	poolBtripFlightModifyFlightInfoRs.Put(v)
}
