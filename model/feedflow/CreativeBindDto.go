package feedflow

import (
	"sync"
)

// CreativeBindDto 结构体
type CreativeBindDto struct {
	// 创意图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 创意名称，同时会展现给客户
	CreativeName string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
	// 创意文案
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 审核状态，W待审核，P审核通过，R审核拒绝
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 审核拒绝原因
	AuditReason string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty"`
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}

var poolCreativeBindDto = sync.Pool{
	New: func() any {
		return new(CreativeBindDto)
	},
}

// GetCreativeBindDto() 从对象池中获取CreativeBindDto
func GetCreativeBindDto() *CreativeBindDto {
	return poolCreativeBindDto.Get().(*CreativeBindDto)
}

// ReleaseCreativeBindDto 释放CreativeBindDto
func ReleaseCreativeBindDto(v *CreativeBindDto) {
	v.ImgUrl = ""
	v.CreativeName = ""
	v.Title = ""
	v.AuditStatus = ""
	v.AuditReason = ""
	v.CreativeId = 0
	v.CampaignId = 0
	v.AdgroupId = 0
	poolCreativeBindDto.Put(v)
}
