package cloudgame

// ShutdownRoomRequest 结构体
type ShutdownRoomRequest struct {
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间ID
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}
