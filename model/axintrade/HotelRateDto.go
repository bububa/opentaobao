package axintrade

// HotelRateDto 结构体
type HotelRateDto struct {
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
}
