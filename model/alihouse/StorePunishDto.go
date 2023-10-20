package alihouse

import (
	"sync"
)

// StorePunishDto 结构体
type StorePunishDto struct {
	// dto
	PunishDtos []PunishDto `json:"punish_dtos,omitempty" xml:"punish_dtos>punish_dto,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolStorePunishDto = sync.Pool{
	New: func() any {
		return new(StorePunishDto)
	},
}

// GetStorePunishDto() 从对象池中获取StorePunishDto
func GetStorePunishDto() *StorePunishDto {
	return poolStorePunishDto.Get().(*StorePunishDto)
}

// ReleaseStorePunishDto 释放StorePunishDto
func ReleaseStorePunishDto(v *StorePunishDto) {
	v.PunishDtos = v.PunishDtos[:0]
	v.StoreId = 0
	poolStorePunishDto.Put(v)
}
