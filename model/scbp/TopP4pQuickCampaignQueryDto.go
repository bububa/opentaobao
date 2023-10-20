package scbp

import (
	"sync"
)

// TopP4pQuickCampaignQueryDto 结构体
type TopP4pQuickCampaignQueryDto struct {
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
	// 每页返回数量
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
}

var poolTopP4pQuickCampaignQueryDto = sync.Pool{
	New: func() any {
		return new(TopP4pQuickCampaignQueryDto)
	},
}

// GetTopP4pQuickCampaignQueryDto() 从对象池中获取TopP4pQuickCampaignQueryDto
func GetTopP4pQuickCampaignQueryDto() *TopP4pQuickCampaignQueryDto {
	return poolTopP4pQuickCampaignQueryDto.Get().(*TopP4pQuickCampaignQueryDto)
}

// ReleaseTopP4pQuickCampaignQueryDto 释放TopP4pQuickCampaignQueryDto
func ReleaseTopP4pQuickCampaignQueryDto(v *TopP4pQuickCampaignQueryDto) {
	v.ToPage = 0
	v.PerPageSize = 0
	poolTopP4pQuickCampaignQueryDto.Put(v)
}
