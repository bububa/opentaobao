package cloudgame

import (
	"sync"
)

// StartGameRequest 结构体
type StartGameRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 主播用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolStartGameRequest = sync.Pool{
	New: func() any {
		return new(StartGameRequest)
	},
}

// GetStartGameRequest() 从对象池中获取StartGameRequest
func GetStartGameRequest() *StartGameRequest {
	return poolStartGameRequest.Get().(*StartGameRequest)
}

// ReleaseStartGameRequest 释放StartGameRequest
func ReleaseStartGameRequest(v *StartGameRequest) {
	v.MixGameId = ""
	v.UserId = ""
	v.Token = ""
	v.RoomId = 0
	poolStartGameRequest.Put(v)
}
