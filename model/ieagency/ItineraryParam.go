package ieagency

// ItineraryParam 结构体
type ItineraryParam struct {
	// 航段信息
	FlightSegmentParams []FlightSegmentParam `json:"flight_segment_params,omitempty" xml:"flight_segment_params>flight_segment_param,omitempty"`
	// 航程序号（从1开始）
	ItineraryRph int64 `json:"itinerary_rph,omitempty" xml:"itinerary_rph,omitempty"`
}
