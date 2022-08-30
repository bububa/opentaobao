package axintrade

// HotelArrivalTime 结构体
type HotelArrivalTime struct {
	// 最晚到店时间
	LatestArrivalTime string `json:"latest_arrival_time,omitempty" xml:"latest_arrival_time,omitempty"`
	// 最早到店时间
	EarliestArrivalTime string `json:"earliest_arrival_time,omitempty" xml:"earliest_arrival_time,omitempty"`
}
