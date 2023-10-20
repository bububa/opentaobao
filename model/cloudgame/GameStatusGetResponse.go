package cloudgame

import (
	"sync"
)

// GameStatusGetResponse 结构体
type GameStatusGetResponse struct {
	// 玩家列表
	PlayerList []OpenGamePlayerDto `json:"player_list,omitempty" xml:"player_list>open_game_player_dto,omitempty"`
	// 游戏会话id
	GameSession string `json:"game_session,omitempty" xml:"game_session,omitempty"`
	// 用户状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 状态详情数据
	StatusData string `json:"status_data,omitempty" xml:"status_data,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 游戏详情
	Game *OpenGameDto `json:"game,omitempty" xml:"game,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolGameStatusGetResponse = sync.Pool{
	New: func() any {
		return new(GameStatusGetResponse)
	},
}

// GetGameStatusGetResponse() 从对象池中获取GameStatusGetResponse
func GetGameStatusGetResponse() *GameStatusGetResponse {
	return poolGameStatusGetResponse.Get().(*GameStatusGetResponse)
}

// ReleaseGameStatusGetResponse 释放GameStatusGetResponse
func ReleaseGameStatusGetResponse(v *GameStatusGetResponse) {
	v.PlayerList = v.PlayerList[:0]
	v.GameSession = ""
	v.Status = ""
	v.ExtInfo = ""
	v.StatusData = ""
	v.UserId = ""
	v.Game = nil
	v.RoomId = 0
	poolGameStatusGetResponse.Put(v)
}
