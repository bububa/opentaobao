package happytrip

// SyncHotelBookingDataRequestDto 结构体
type SyncHotelBookingDataRequestDto struct {
	// 预订信息数据
	Segments []Null `json:"segments,omitempty" xml:"segments>null,omitempty"`
}
