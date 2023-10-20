package jym

import (
	"sync"
)

// SteamShadowDto 结构体
type SteamShadowDto struct {
	// 影子ID
	ShadowUid string `json:"shadow_uid,omitempty" xml:"shadow_uid,omitempty"`
}

var poolSteamShadowDto = sync.Pool{
	New: func() any {
		return new(SteamShadowDto)
	},
}

// GetSteamShadowDto() 从对象池中获取SteamShadowDto
func GetSteamShadowDto() *SteamShadowDto {
	return poolSteamShadowDto.Get().(*SteamShadowDto)
}

// ReleaseSteamShadowDto 释放SteamShadowDto
func ReleaseSteamShadowDto(v *SteamShadowDto) {
	v.ShadowUid = ""
	poolSteamShadowDto.Put(v)
}
