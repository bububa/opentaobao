package alihouse

import (
	"sync"
)

// PunishDto 结构体
type PunishDto struct {
	// 处罚级别
	PunishLevel int64 `json:"punish_level,omitempty" xml:"punish_level,omitempty"`
	// 业务类型
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 处罚类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 类目编码
	CategoryCode int64 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 是否处罚
	IsPunish bool `json:"is_punish,omitempty" xml:"is_punish,omitempty"`
}

var poolPunishDto = sync.Pool{
	New: func() any {
		return new(PunishDto)
	},
}

// GetPunishDto() 从对象池中获取PunishDto
func GetPunishDto() *PunishDto {
	return poolPunishDto.Get().(*PunishDto)
}

// ReleasePunishDto 释放PunishDto
func ReleasePunishDto(v *PunishDto) {
	v.PunishLevel = 0
	v.BusinessType = 0
	v.Type = 0
	v.Version = 0
	v.CategoryCode = 0
	v.IsPunish = false
	poolPunishDto.Put(v)
}
