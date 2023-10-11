package jym

// JymMaxPriceRequestDto 结构体
type JymMaxPriceRequestDto struct {
	// 交易猫游戏ID
	JymGameId string `json:"jym_game_id,omitempty" xml:"jym_game_id,omitempty"`
	// 渠道标识
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}
