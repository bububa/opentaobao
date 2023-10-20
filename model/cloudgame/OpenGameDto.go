package cloudgame

import (
	"sync"
)

// OpenGameDto 结构体
type OpenGameDto struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 游戏名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建者id
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 玩家id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 位置索引
	PlayerIndex string `json:"player_index,omitempty" xml:"player_index,omitempty"`
	// 游戏总人数
	TotalPlayerNum int64 `json:"total_player_num,omitempty" xml:"total_player_num,omitempty"`
	// 游戏当前人数
	PlayerNum int64 `json:"player_num,omitempty" xml:"player_num,omitempty"`
}

var poolOpenGameDto = sync.Pool{
	New: func() any {
		return new(OpenGameDto)
	},
}

// GetOpenGameDto() 从对象池中获取OpenGameDto
func GetOpenGameDto() *OpenGameDto {
	return poolOpenGameDto.Get().(*OpenGameDto)
}

// ReleaseOpenGameDto 释放OpenGameDto
func ReleaseOpenGameDto(v *OpenGameDto) {
	v.MixGameId = ""
	v.Name = ""
	v.Creator = ""
	v.UserId = ""
	v.PlayerIndex = ""
	v.TotalPlayerNum = 0
	v.PlayerNum = 0
	poolOpenGameDto.Put(v)
}
