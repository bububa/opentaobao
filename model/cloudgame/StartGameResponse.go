package cloudgame

import (
	"sync"
)

// StartGameResponse 结构体
type StartGameResponse struct {
	// 玩家列表
	PlayerList []OpenGamePlayerDto `json:"player_list,omitempty" xml:"player_list>open_game_player_dto,omitempty"`
	// 游戏会话id
	GameSession string `json:"game_session,omitempty" xml:"game_session,omitempty"`
	// 联机码
	JoinCode string `json:"join_code,omitempty" xml:"join_code,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 游戏详情
	Game *OpenGameDto `json:"game,omitempty" xml:"game,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 联机信息
	Slot *GetSlotResponse `json:"slot,omitempty" xml:"slot,omitempty"`
}

var poolStartGameResponse = sync.Pool{
	New: func() any {
		return new(StartGameResponse)
	},
}

// GetStartGameResponse() 从对象池中获取StartGameResponse
func GetStartGameResponse() *StartGameResponse {
	return poolStartGameResponse.Get().(*StartGameResponse)
}

// ReleaseStartGameResponse 释放StartGameResponse
func ReleaseStartGameResponse(v *StartGameResponse) {
	v.PlayerList = v.PlayerList[:0]
	v.GameSession = ""
	v.JoinCode = ""
	v.ExtInfo = ""
	v.Game = nil
	v.RoomId = 0
	v.Slot = nil
	poolStartGameResponse.Put(v)
}
