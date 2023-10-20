package cloudgame

import (
	"sync"
)

// StopUserGameResponse 结构体
type StopUserGameResponse struct {
	// 玩家列表
	PlayerList []OpenGamePlayerDto `json:"player_list,omitempty" xml:"player_list>open_game_player_dto,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 游戏详情
	Game *OpenGameDto `json:"game,omitempty" xml:"game,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolStopUserGameResponse = sync.Pool{
	New: func() any {
		return new(StopUserGameResponse)
	},
}

// GetStopUserGameResponse() 从对象池中获取StopUserGameResponse
func GetStopUserGameResponse() *StopUserGameResponse {
	return poolStopUserGameResponse.Get().(*StopUserGameResponse)
}

// ReleaseStopUserGameResponse 释放StopUserGameResponse
func ReleaseStopUserGameResponse(v *StopUserGameResponse) {
	v.PlayerList = v.PlayerList[:0]
	v.ExtInfo = ""
	v.Game = nil
	v.RoomId = 0
	poolStopUserGameResponse.Put(v)
}
