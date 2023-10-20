package cloudgame

import (
	"sync"
)

// OpenGamePlayerDto 结构体
type OpenGamePlayerDto struct {
	// 玩家id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 位置索引
	PlayerIndex string `json:"player_index,omitempty" xml:"player_index,omitempty"`
}

var poolOpenGamePlayerDto = sync.Pool{
	New: func() any {
		return new(OpenGamePlayerDto)
	},
}

// GetOpenGamePlayerDto() 从对象池中获取OpenGamePlayerDto
func GetOpenGamePlayerDto() *OpenGamePlayerDto {
	return poolOpenGamePlayerDto.Get().(*OpenGamePlayerDto)
}

// ReleaseOpenGamePlayerDto 释放OpenGamePlayerDto
func ReleaseOpenGamePlayerDto(v *OpenGamePlayerDto) {
	v.UserId = ""
	v.PlayerIndex = ""
	poolOpenGamePlayerDto.Put(v)
}
