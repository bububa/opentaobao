package btrip

import (
	"sync"
)

// BtripFlightModifySearchPriceRs 结构体
type BtripFlightModifySearchPriceRs struct {
	// 航班列表
	FlightInfoList []FlightInfoDto `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info_dto,omitempty"`
}

var poolBtripFlightModifySearchPriceRs = sync.Pool{
	New: func() any {
		return new(BtripFlightModifySearchPriceRs)
	},
}

// GetBtripFlightModifySearchPriceRs() 从对象池中获取BtripFlightModifySearchPriceRs
func GetBtripFlightModifySearchPriceRs() *BtripFlightModifySearchPriceRs {
	return poolBtripFlightModifySearchPriceRs.Get().(*BtripFlightModifySearchPriceRs)
}

// ReleaseBtripFlightModifySearchPriceRs 释放BtripFlightModifySearchPriceRs
func ReleaseBtripFlightModifySearchPriceRs(v *BtripFlightModifySearchPriceRs) {
	v.FlightInfoList = v.FlightInfoList[:0]
	poolBtripFlightModifySearchPriceRs.Put(v)
}
