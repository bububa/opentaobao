package btrip

// BtripFlightSearchListRs 结构体
type BtripFlightSearchListRs struct {
	// 航班列表
	FlightInfoList []FlightInfoDto `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info_dto,omitempty"`
	// 是否可更换PNR出票
	IsReplacePnr bool `json:"is_replace_pnr,omitempty" xml:"is_replace_pnr,omitempty"`
}
