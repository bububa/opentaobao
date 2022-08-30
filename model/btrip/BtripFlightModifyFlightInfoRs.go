package btrip

// BtripFlightModifyFlightInfoRs 结构体
type BtripFlightModifyFlightInfoRs struct {
	// 改签航班信息列表
	FlightInfoList []FlightInfoDto `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info_dto,omitempty"`
}
