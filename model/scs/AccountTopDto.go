package scs

import (
	"sync"
)

// AccountTopDto 结构体
type AccountTopDto struct {
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolAccountTopDto = sync.Pool{
	New: func() any {
		return new(AccountTopDto)
	},
}

// GetAccountTopDto() 从对象池中获取AccountTopDto
func GetAccountTopDto() *AccountTopDto {
	return poolAccountTopDto.Get().(*AccountTopDto)
}

// ReleaseAccountTopDto 释放AccountTopDto
func ReleaseAccountTopDto(v *AccountTopDto) {
	v.CampaignId = 0
	poolAccountTopDto.Put(v)
}
