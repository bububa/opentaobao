package jym

import (
	"sync"
)

// OutSideQueryGamePropertyInfoRequestDto 结构体
type OutSideQueryGamePropertyInfoRequestDto struct {
	// 交易猫游戏id
	GameId string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

var poolOutSideQueryGamePropertyInfoRequestDto = sync.Pool{
	New: func() any {
		return new(OutSideQueryGamePropertyInfoRequestDto)
	},
}

// GetOutSideQueryGamePropertyInfoRequestDto() 从对象池中获取OutSideQueryGamePropertyInfoRequestDto
func GetOutSideQueryGamePropertyInfoRequestDto() *OutSideQueryGamePropertyInfoRequestDto {
	return poolOutSideQueryGamePropertyInfoRequestDto.Get().(*OutSideQueryGamePropertyInfoRequestDto)
}

// ReleaseOutSideQueryGamePropertyInfoRequestDto 释放OutSideQueryGamePropertyInfoRequestDto
func ReleaseOutSideQueryGamePropertyInfoRequestDto(v *OutSideQueryGamePropertyInfoRequestDto) {
	v.GameId = ""
	poolOutSideQueryGamePropertyInfoRequestDto.Put(v)
}
