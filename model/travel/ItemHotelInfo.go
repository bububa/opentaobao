package travel

// ItemHotelInfo 
type ItemHotelInfo struct {
    // 结构化酒店信息 酒店结构化信息和文本描述二选一
    HotelList   []ItemHotel `json:"hotel_list,omitempty" xml:"hotel_list>item_hotel,omitempty"`
}
