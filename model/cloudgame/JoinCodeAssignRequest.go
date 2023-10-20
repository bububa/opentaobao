package cloudgame

import (
	"sync"
)

// JoinCodeAssignRequest 结构体
type JoinCodeAssignRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 主播用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 玩家位置
	PlayIndex string `json:"play_index,omitempty" xml:"play_index,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolJoinCodeAssignRequest = sync.Pool{
	New: func() any {
		return new(JoinCodeAssignRequest)
	},
}

// GetJoinCodeAssignRequest() 从对象池中获取JoinCodeAssignRequest
func GetJoinCodeAssignRequest() *JoinCodeAssignRequest {
	return poolJoinCodeAssignRequest.Get().(*JoinCodeAssignRequest)
}

// ReleaseJoinCodeAssignRequest 释放JoinCodeAssignRequest
func ReleaseJoinCodeAssignRequest(v *JoinCodeAssignRequest) {
	v.MixGameId = ""
	v.UserId = ""
	v.Token = ""
	v.PlayIndex = ""
	v.RoomId = 0
	poolJoinCodeAssignRequest.Put(v)
}
