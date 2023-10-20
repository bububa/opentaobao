package cloudgame

import (
	"sync"
)

// HeartBeatRequest 结构体
type HeartBeatRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 主播用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 心跳序列
	Seq int64 `json:"seq,omitempty" xml:"seq,omitempty"`
	// 心跳间隔(s)
	Interval int64 `json:"interval,omitempty" xml:"interval,omitempty"`
}

var poolHeartBeatRequest = sync.Pool{
	New: func() any {
		return new(HeartBeatRequest)
	},
}

// GetHeartBeatRequest() 从对象池中获取HeartBeatRequest
func GetHeartBeatRequest() *HeartBeatRequest {
	return poolHeartBeatRequest.Get().(*HeartBeatRequest)
}

// ReleaseHeartBeatRequest 释放HeartBeatRequest
func ReleaseHeartBeatRequest(v *HeartBeatRequest) {
	v.MixGameId = ""
	v.UserId = ""
	v.Token = ""
	v.RoomId = 0
	v.Seq = 0
	v.Interval = 0
	poolHeartBeatRequest.Put(v)
}
