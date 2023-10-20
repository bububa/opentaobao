package feedflow

import (
	"sync"
)

// CreativeBindQueryDto 结构体
type CreativeBindQueryDto struct {
	// 创意id列表
	CreativeIdList []int64 `json:"creative_id_list,omitempty" xml:"creative_id_list>int64,omitempty"`
	// 创意名称
	CreativeName string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
	// 审核状态，W待审核，P审核通过，R审核拒绝
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 分页页码
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页起始位置
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolCreativeBindQueryDto = sync.Pool{
	New: func() any {
		return new(CreativeBindQueryDto)
	},
}

// GetCreativeBindQueryDto() 从对象池中获取CreativeBindQueryDto
func GetCreativeBindQueryDto() *CreativeBindQueryDto {
	return poolCreativeBindQueryDto.Get().(*CreativeBindQueryDto)
}

// ReleaseCreativeBindQueryDto 释放CreativeBindQueryDto
func ReleaseCreativeBindQueryDto(v *CreativeBindQueryDto) {
	v.CreativeIdList = v.CreativeIdList[:0]
	v.CreativeName = ""
	v.AuditStatus = ""
	v.AdgroupId = 0
	v.PageSize = 0
	v.Offset = 0
	v.CampaignId = 0
	poolCreativeBindQueryDto.Put(v)
}
