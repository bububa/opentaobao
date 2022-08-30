package happytrip

// SyncHotelBookingDataRequestDto 结构体
type SyncHotelBookingDataRequestDto struct {
	// 预订信息数据
	Segments []string `json:"segments,omitempty" xml:"segments>string,omitempty"`
}
