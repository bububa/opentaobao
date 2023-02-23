package bus

// MerchantBusLineInfo 结构体
type MerchantBusLineInfo struct {
	// 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
	// 到达站点
	ArriveStationName string `json:"arrive_station_name,omitempty" xml:"arrive_station_name,omitempty"`
	// 出发日期
	DepartDate string `json:"depart_date,omitempty" xml:"depart_date,omitempty"`
	// 车牌号
	BusNumber string `json:"bus_number,omitempty" xml:"bus_number,omitempty"`
	// 出发时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 班次id
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
	// 出发站名称
	StartStationName string `json:"start_station_name,omitempty" xml:"start_station_name,omitempty"`
}
