package feedflow

import (
	"sync"
)

// RptQueryDto 结构体
type RptQueryDto struct {
	// 查询开始日期
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 查询结束日期
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 查询日期
	LogDate string `json:"log_date,omitempty" xml:"log_date,omitempty"`
	// 归因日期（有效值为3、7、15、30）
	Effect int64 `json:"effect,omitempty" xml:"effect,omitempty"`
	// 结束小时
	EndHourId int64 `json:"end_hour_id,omitempty" xml:"end_hour_id,omitempty"`
	// 开始小时
	StartHourId int64 `json:"start_hour_id,omitempty" xml:"start_hour_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 资源位id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 定向id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
}

var poolRptQueryDto = sync.Pool{
	New: func() any {
		return new(RptQueryDto)
	},
}

// GetRptQueryDto() 从对象池中获取RptQueryDto
func GetRptQueryDto() *RptQueryDto {
	return poolRptQueryDto.Get().(*RptQueryDto)
}

// ReleaseRptQueryDto 释放RptQueryDto
func ReleaseRptQueryDto(v *RptQueryDto) {
	v.StartTime = ""
	v.EndTime = ""
	v.LogDate = ""
	v.Effect = 0
	v.EndHourId = 0
	v.StartHourId = 0
	v.CampaignId = 0
	v.AdgroupId = 0
	v.AdzoneId = 0
	v.CreativeId = 0
	v.CrowdId = 0
	poolRptQueryDto.Put(v)
}
