package jym

import (
	"sync"
)

// JymMaxPriceOrderInfoDto 结构体
type JymMaxPriceOrderInfoDto struct {
	// 交易猫游戏ID
	JymGameId string `json:"jym_game_id,omitempty" xml:"jym_game_id,omitempty"`
	// 交易猫游戏名称
	JymGameName string `json:"jym_game_name,omitempty" xml:"jym_game_name,omitempty"`
	// 订单价格
	OrderPrice int64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
}

var poolJymMaxPriceOrderInfoDto = sync.Pool{
	New: func() any {
		return new(JymMaxPriceOrderInfoDto)
	},
}

// GetJymMaxPriceOrderInfoDto() 从对象池中获取JymMaxPriceOrderInfoDto
func GetJymMaxPriceOrderInfoDto() *JymMaxPriceOrderInfoDto {
	return poolJymMaxPriceOrderInfoDto.Get().(*JymMaxPriceOrderInfoDto)
}

// ReleaseJymMaxPriceOrderInfoDto 释放JymMaxPriceOrderInfoDto
func ReleaseJymMaxPriceOrderInfoDto(v *JymMaxPriceOrderInfoDto) {
	v.JymGameId = ""
	v.JymGameName = ""
	v.OrderPrice = 0
	poolJymMaxPriceOrderInfoDto.Put(v)
}
