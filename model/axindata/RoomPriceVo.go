package axindata

// RoomPriceVo 结构体
type RoomPriceVo struct {
	// 价库信息
	RateList []RateVo `json:"rate_list,omitempty" xml:"rate_list>rate_vo,omitempty"`
	// 房型信息
	StdRoomInfo *StdRoomType `json:"std_room_info,omitempty" xml:"std_room_info,omitempty"`
}
