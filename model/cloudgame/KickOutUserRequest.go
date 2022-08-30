package cloudgame

// KickOutUserRequest 结构体
type KickOutUserRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 踢出用户id
	KickOutUserId string `json:"kick_out_user_id,omitempty" xml:"kick_out_user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}
