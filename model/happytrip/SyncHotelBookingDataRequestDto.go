package happytrip

// SyncHotelBookingDataRequestDto 
type SyncHotelBookingDataRequestDto struct {
    // 预订信息数据
    Segments   []Null `json:"segments,omitempty" xml:"segments>null,omitempty"`
}
