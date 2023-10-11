package jym

// OutSideSyncGamePropertyRequestDto 结构体
type OutSideSyncGamePropertyRequestDto struct {
	// 交易猫游戏id
	GameId string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	// 压缩加密后的属性信息
	GamePropertiesCompressed string `json:"game_properties_compressed,omitempty" xml:"game_properties_compressed,omitempty"`
}
