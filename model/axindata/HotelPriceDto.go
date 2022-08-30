package axindata

// HotelPriceDto 结构体
type HotelPriceDto struct {
	// 房型信息列表
	RoomList []RoomPriceDto `json:"room_list,omitempty" xml:"room_list>room_price_dto,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
