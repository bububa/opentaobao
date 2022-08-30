package cloudgame

// CreateRoomResponse 结构体
type CreateRoomResponse struct {
	// 扩展数据
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}
