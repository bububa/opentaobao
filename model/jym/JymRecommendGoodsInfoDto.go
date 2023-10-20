package jym

import (
	"sync"
)

// JymRecommendGoodsInfoDto 结构体
type JymRecommendGoodsInfoDto struct {
	// 推荐商品
	GoodsList []JymSingleGoodsDto `json:"goods_list,omitempty" xml:"goods_list>jym_single_goods_dto,omitempty"`
	// 交易猫游戏ID
	JymGameId string `json:"jym_game_id,omitempty" xml:"jym_game_id,omitempty"`
}

var poolJymRecommendGoodsInfoDto = sync.Pool{
	New: func() any {
		return new(JymRecommendGoodsInfoDto)
	},
}

// GetJymRecommendGoodsInfoDto() 从对象池中获取JymRecommendGoodsInfoDto
func GetJymRecommendGoodsInfoDto() *JymRecommendGoodsInfoDto {
	return poolJymRecommendGoodsInfoDto.Get().(*JymRecommendGoodsInfoDto)
}

// ReleaseJymRecommendGoodsInfoDto 释放JymRecommendGoodsInfoDto
func ReleaseJymRecommendGoodsInfoDto(v *JymRecommendGoodsInfoDto) {
	v.GoodsList = v.GoodsList[:0]
	v.JymGameId = ""
	poolJymRecommendGoodsInfoDto.Put(v)
}
