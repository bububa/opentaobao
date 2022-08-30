package bus

// MerchantBusLineInfo 结构体
type MerchantBusLineInfo struct {
	// 车次编号
	BusNumber string `json:"bus_number,omitempty" xml:"bus_number,omitempty"`
	// 预计到达时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 出发时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 发车时间
	DepardDate string `json:"depard_date,omitempty" xml:"depard_date,omitempty"`
	// 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
	// 终点站
	ArriveStationName string `json:"arrive_station_name,omitempty" xml:"arrive_station_name,omitempty"`
	// 目的城市
	EndCityName string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty"`
	// 出发站地址
	StartStationAddress string `json:"start_station_address,omitempty" xml:"start_station_address,omitempty"`
	// 出发站
	StartStationName string `json:"start_station_name,omitempty" xml:"start_station_name,omitempty"`
	// 出发城市
	StartCityName string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
	// 商家路线id
	AgentRouteId int64 `json:"agent_route_id,omitempty" xml:"agent_route_id,omitempty"`
}
