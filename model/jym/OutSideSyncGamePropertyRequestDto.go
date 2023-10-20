package jym

import (
	"sync"
)

// OutSideSyncGamePropertyRequestDto 结构体
type OutSideSyncGamePropertyRequestDto struct {
	// 交易猫游戏id
	GameId string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	// 压缩加密后的属性信息
	GamePropertiesCompressed string `json:"game_properties_compressed,omitempty" xml:"game_properties_compressed,omitempty"`
}

var poolOutSideSyncGamePropertyRequestDto = sync.Pool{
	New: func() any {
		return new(OutSideSyncGamePropertyRequestDto)
	},
}

// GetOutSideSyncGamePropertyRequestDto() 从对象池中获取OutSideSyncGamePropertyRequestDto
func GetOutSideSyncGamePropertyRequestDto() *OutSideSyncGamePropertyRequestDto {
	return poolOutSideSyncGamePropertyRequestDto.Get().(*OutSideSyncGamePropertyRequestDto)
}

// ReleaseOutSideSyncGamePropertyRequestDto 释放OutSideSyncGamePropertyRequestDto
func ReleaseOutSideSyncGamePropertyRequestDto(v *OutSideSyncGamePropertyRequestDto) {
	v.GameId = ""
	v.GamePropertiesCompressed = ""
	poolOutSideSyncGamePropertyRequestDto.Put(v)
}
