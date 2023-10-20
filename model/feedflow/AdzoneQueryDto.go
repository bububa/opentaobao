package feedflow

import (
	"sync"
)

// AdzoneQueryDto 结构体
type AdzoneQueryDto struct {
	// 广告位id列表
	AdzoneIdList []int64 `json:"adzone_id_list,omitempty" xml:"adzone_id_list>int64,omitempty"`
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolAdzoneQueryDto = sync.Pool{
	New: func() any {
		return new(AdzoneQueryDto)
	},
}

// GetAdzoneQueryDto() 从对象池中获取AdzoneQueryDto
func GetAdzoneQueryDto() *AdzoneQueryDto {
	return poolAdzoneQueryDto.Get().(*AdzoneQueryDto)
}

// ReleaseAdzoneQueryDto 释放AdzoneQueryDto
func ReleaseAdzoneQueryDto(v *AdzoneQueryDto) {
	v.AdzoneIdList = v.AdzoneIdList[:0]
	v.AdzoneName = ""
	v.CampaignId = 0
	poolAdzoneQueryDto.Put(v)
}
