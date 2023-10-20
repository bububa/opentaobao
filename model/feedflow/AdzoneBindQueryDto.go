package feedflow

import (
	"sync"
)

// AdzoneBindQueryDto 结构体
type AdzoneBindQueryDto struct {
	// 广告位id列表
	AdzoneIdList []int64 `json:"adzone_id_list,omitempty" xml:"adzone_id_list>int64,omitempty"`
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 分页条件
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页条件
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolAdzoneBindQueryDto = sync.Pool{
	New: func() any {
		return new(AdzoneBindQueryDto)
	},
}

// GetAdzoneBindQueryDto() 从对象池中获取AdzoneBindQueryDto
func GetAdzoneBindQueryDto() *AdzoneBindQueryDto {
	return poolAdzoneBindQueryDto.Get().(*AdzoneBindQueryDto)
}

// ReleaseAdzoneBindQueryDto 释放AdzoneBindQueryDto
func ReleaseAdzoneBindQueryDto(v *AdzoneBindQueryDto) {
	v.AdzoneIdList = v.AdzoneIdList[:0]
	v.AdzoneName = ""
	v.AdgroupId = 0
	v.PageSize = 0
	v.Offset = 0
	v.CampaignId = 0
	poolAdzoneBindQueryDto.Put(v)
}
