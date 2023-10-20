package cloudgame

import (
	"sync"
)

// GameStatusGetRequest 结构体
type GameStatusGetRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolGameStatusGetRequest = sync.Pool{
	New: func() any {
		return new(GameStatusGetRequest)
	},
}

// GetGameStatusGetRequest() 从对象池中获取GameStatusGetRequest
func GetGameStatusGetRequest() *GameStatusGetRequest {
	return poolGameStatusGetRequest.Get().(*GameStatusGetRequest)
}

// ReleaseGameStatusGetRequest 释放GameStatusGetRequest
func ReleaseGameStatusGetRequest(v *GameStatusGetRequest) {
	v.MixGameId = ""
	v.Token = ""
	v.RoomId = 0
	poolGameStatusGetRequest.Put(v)
}
