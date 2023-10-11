package car

// TransferFlightInfo 结构体
type TransferFlightInfo struct {
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 接机到达/送机出发航站楼
	AirportTerminal string `json:"airport_terminal,omitempty" xml:"airport_terminal,omitempty"`
	// 出发 机场三字码
	DepAirPortCode string `json:"dep_air_port_code,omitempty" xml:"dep_air_port_code,omitempty"`
	// 航班到达时间
	FlightArrivedDate string `json:"flight_arrived_date,omitempty" xml:"flight_arrived_date,omitempty"`
	// 到达 机场三字码
	DesAirPortCode string `json:"des_air_port_code,omitempty" xml:"des_air_port_code,omitempty"`
	// 航班起飞日期
	FlightFlyDate string `json:"flight_fly_date,omitempty" xml:"flight_fly_date,omitempty"`
}
