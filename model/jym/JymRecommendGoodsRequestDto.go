package jym

import (
	"sync"
)

// JymRecommendGoodsRequestDto 结构体
type JymRecommendGoodsRequestDto struct {
	// 交易猫游戏ID
	JymGameId string `json:"jym_game_id,omitempty" xml:"jym_game_id,omitempty"`
	// 渠道标识
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 最低价
	MinPrice int64 `json:"min_price,omitempty" xml:"min_price,omitempty"`
	// 最高价
	MaxPrice int64 `json:"max_price,omitempty" xml:"max_price,omitempty"`
}

var poolJymRecommendGoodsRequestDto = sync.Pool{
	New: func() any {
		return new(JymRecommendGoodsRequestDto)
	},
}

// GetJymRecommendGoodsRequestDto() 从对象池中获取JymRecommendGoodsRequestDto
func GetJymRecommendGoodsRequestDto() *JymRecommendGoodsRequestDto {
	return poolJymRecommendGoodsRequestDto.Get().(*JymRecommendGoodsRequestDto)
}

// ReleaseJymRecommendGoodsRequestDto 释放JymRecommendGoodsRequestDto
func ReleaseJymRecommendGoodsRequestDto(v *JymRecommendGoodsRequestDto) {
	v.JymGameId = ""
	v.Channel = ""
	v.MinPrice = 0
	v.MaxPrice = 0
	poolJymRecommendGoodsRequestDto.Put(v)
}
