package btrip

// PassengerSegmentInfo 结构体
type PassengerSegmentInfo struct {
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 用户编号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
