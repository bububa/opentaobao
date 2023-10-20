package cloudgame

import (
	"sync"
)

// UserGameStatusGetRequest 结构体
type UserGameStatusGetRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 主播用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolUserGameStatusGetRequest = sync.Pool{
	New: func() any {
		return new(UserGameStatusGetRequest)
	},
}

// GetUserGameStatusGetRequest() 从对象池中获取UserGameStatusGetRequest
func GetUserGameStatusGetRequest() *UserGameStatusGetRequest {
	return poolUserGameStatusGetRequest.Get().(*UserGameStatusGetRequest)
}

// ReleaseUserGameStatusGetRequest 释放UserGameStatusGetRequest
func ReleaseUserGameStatusGetRequest(v *UserGameStatusGetRequest) {
	v.MixGameId = ""
	v.UserId = ""
	v.Token = ""
	v.RoomId = 0
	poolUserGameStatusGetRequest.Put(v)
}
