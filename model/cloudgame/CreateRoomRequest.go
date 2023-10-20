package cloudgame

import (
	"sync"
)

// CreateRoomRequest 结构体
type CreateRoomRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 主播用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
}

var poolCreateRoomRequest = sync.Pool{
	New: func() any {
		return new(CreateRoomRequest)
	},
}

// GetCreateRoomRequest() 从对象池中获取CreateRoomRequest
func GetCreateRoomRequest() *CreateRoomRequest {
	return poolCreateRoomRequest.Get().(*CreateRoomRequest)
}

// ReleaseCreateRoomRequest 释放CreateRoomRequest
func ReleaseCreateRoomRequest(v *CreateRoomRequest) {
	v.MixGameId = ""
	v.UserId = ""
	v.Token = ""
	poolCreateRoomRequest.Put(v)
}
