package tmallservice

// Reservation 结构体
type Reservation struct {
	// 预约时间
	ReserveTimeEnd string `json:"reserve_time_end,omitempty" xml:"reserve_time_end,omitempty"`
	// 预约时间
	ReserveTimeStart string `json:"reserve_time_start,omitempty" xml:"reserve_time_start,omitempty"`
	// 预约来源
	ReserveRequestSourceType string `json:"reserve_request_source_type,omitempty" xml:"reserve_request_source_type,omitempty"`
}
