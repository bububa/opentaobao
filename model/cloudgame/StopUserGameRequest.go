package cloudgame

import (
	"sync"
)

// StopUserGameRequest 结构体
type StopUserGameRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolStopUserGameRequest = sync.Pool{
	New: func() any {
		return new(StopUserGameRequest)
	},
}

// GetStopUserGameRequest() 从对象池中获取StopUserGameRequest
func GetStopUserGameRequest() *StopUserGameRequest {
	return poolStopUserGameRequest.Get().(*StopUserGameRequest)
}

// ReleaseStopUserGameRequest 释放StopUserGameRequest
func ReleaseStopUserGameRequest(v *StopUserGameRequest) {
	v.MixGameId = ""
	v.UserId = ""
	v.Token = ""
	v.RoomId = 0
	poolStopUserGameRequest.Put(v)
}
