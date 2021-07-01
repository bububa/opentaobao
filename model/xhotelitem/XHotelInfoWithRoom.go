package xhotelitem

// XHotelInfoWithRoom 结构体
type XHotelInfoWithRoom struct {
	// 房型基础信息
	RoomTypeList []RoomType `json:"room_type_list,omitempty" xml:"room_type_list>room_type,omitempty"`
}
