package btrip

// BtripFlightModifySearchPriceRs 结构体
type BtripFlightModifySearchPriceRs struct {
	// 航班列表
	FlightInfoList []FlightInfoDto `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info_dto,omitempty"`
}
