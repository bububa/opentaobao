package alitripmerchant

// Room 结构体
type Room struct {
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 价格定义名
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 房间面积大小
	RoomArea string `json:"room_area,omitempty" xml:"room_area,omitempty"`
	// 房间床型描述
	BedTypeDesc string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
	// 房型最大入住人数
	MaxCheckInNumber int64 `json:"max_check_in_number,omitempty" xml:"max_check_in_number,omitempty"`
}
