package jym

import (
	"sync"
)

// JymMaxPriceRequestDto 结构体
type JymMaxPriceRequestDto struct {
	// 交易猫游戏ID
	JymGameId string `json:"jym_game_id,omitempty" xml:"jym_game_id,omitempty"`
	// 渠道标识
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}

var poolJymMaxPriceRequestDto = sync.Pool{
	New: func() any {
		return new(JymMaxPriceRequestDto)
	},
}

// GetJymMaxPriceRequestDto() 从对象池中获取JymMaxPriceRequestDto
func GetJymMaxPriceRequestDto() *JymMaxPriceRequestDto {
	return poolJymMaxPriceRequestDto.Get().(*JymMaxPriceRequestDto)
}

// ReleaseJymMaxPriceRequestDto 释放JymMaxPriceRequestDto
func ReleaseJymMaxPriceRequestDto(v *JymMaxPriceRequestDto) {
	v.JymGameId = ""
	v.Channel = ""
	poolJymMaxPriceRequestDto.Put(v)
}
