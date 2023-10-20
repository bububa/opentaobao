package scbp

import (
	"sync"
)

// TopP4pQuickForbiddenWordDto 结构体
type TopP4pQuickForbiddenWordDto struct {
	// 屏蔽词
	ForbiddenWord []string `json:"forbidden_word,omitempty" xml:"forbidden_word>string,omitempty"`
	// 1代表新增屏蔽词，2代表删除屏蔽词
	Action int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 定向推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolTopP4pQuickForbiddenWordDto = sync.Pool{
	New: func() any {
		return new(TopP4pQuickForbiddenWordDto)
	},
}

// GetTopP4pQuickForbiddenWordDto() 从对象池中获取TopP4pQuickForbiddenWordDto
func GetTopP4pQuickForbiddenWordDto() *TopP4pQuickForbiddenWordDto {
	return poolTopP4pQuickForbiddenWordDto.Get().(*TopP4pQuickForbiddenWordDto)
}

// ReleaseTopP4pQuickForbiddenWordDto 释放TopP4pQuickForbiddenWordDto
func ReleaseTopP4pQuickForbiddenWordDto(v *TopP4pQuickForbiddenWordDto) {
	v.ForbiddenWord = v.ForbiddenWord[:0]
	v.Action = 0
	v.CampaignId = 0
	poolTopP4pQuickForbiddenWordDto.Put(v)
}
