package cloudgame

import (
	"sync"
)

// ShutdownRoomResponse 结构体
type ShutdownRoomResponse struct {
	// 游戏当前剩余玩家列表
	PlayerList []OpenGamePlayerDto `json:"player_list,omitempty" xml:"player_list>open_game_player_dto,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 游戏详情
	Game *OpenGameDto `json:"game,omitempty" xml:"game,omitempty"`
	// 房间号
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolShutdownRoomResponse = sync.Pool{
	New: func() any {
		return new(ShutdownRoomResponse)
	},
}

// GetShutdownRoomResponse() 从对象池中获取ShutdownRoomResponse
func GetShutdownRoomResponse() *ShutdownRoomResponse {
	return poolShutdownRoomResponse.Get().(*ShutdownRoomResponse)
}

// ReleaseShutdownRoomResponse 释放ShutdownRoomResponse
func ReleaseShutdownRoomResponse(v *ShutdownRoomResponse) {
	v.PlayerList = v.PlayerList[:0]
	v.ExtInfo = ""
	v.Game = nil
	v.RoomId = 0
	poolShutdownRoomResponse.Put(v)
}
