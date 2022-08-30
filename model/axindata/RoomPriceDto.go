package axindata

// RoomPriceDto 结构体
type RoomPriceDto struct {
	// rate信息列表
	RateList []RateDto `json:"rate_list,omitempty" xml:"rate_list>rate_dto,omitempty"`
	// 房型信息
	StdRoomInfo *StdRoomType `json:"std_room_info,omitempty" xml:"std_room_info,omitempty"`
}
