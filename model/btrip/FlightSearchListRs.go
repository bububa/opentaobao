package btrip

// FlightSearchListRs 结构体
type FlightSearchListRs struct {
	// 航班列表
	FlightList []FlightInfoDto `json:"flight_list,omitempty" xml:"flight_list>flight_info_dto,omitempty"`
	// 是否可更换PNR出票
	IsReplacePnr bool `json:"is_replace_pnr,omitempty" xml:"is_replace_pnr,omitempty"`
}
