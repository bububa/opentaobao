package xhotelitem

// XhotelInfoWithRoom 结构体
type XhotelInfoWithRoom struct {
	// 房型基础信息
	RoomTypeList []RoomType `json:"room_type_list,omitempty" xml:"room_type_list>room_type,omitempty"`
}
