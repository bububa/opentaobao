package nrt

import (
	"sync"
)

// NrtSceneActivityDto 结构体
type NrtSceneActivityDto struct {
	// 权益列表
	BenefitList []NrtBenefitDto `json:"benefit_list,omitempty" xml:"benefit_list>nrt_benefit_dto,omitempty"`
	// 渠道
	ChannelList []ChannelDto `json:"channel_list,omitempty" xml:"channel_list>channel_dto,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 创建者id
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 活动类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 修改人id
	ModifiedId string `json:"modified_id,omitempty" xml:"modified_id,omitempty"`
	// 有价礼包活动对应的商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolNrtSceneActivityDto = sync.Pool{
	New: func() any {
		return new(NrtSceneActivityDto)
	},
}

// GetNrtSceneActivityDto() 从对象池中获取NrtSceneActivityDto
func GetNrtSceneActivityDto() *NrtSceneActivityDto {
	return poolNrtSceneActivityDto.Get().(*NrtSceneActivityDto)
}

// ReleaseNrtSceneActivityDto 释放NrtSceneActivityDto
func ReleaseNrtSceneActivityDto(v *NrtSceneActivityDto) {
	v.BenefitList = v.BenefitList[:0]
	v.ChannelList = v.ChannelList[:0]
	v.BizCode = ""
	v.CreatorId = ""
	v.Extra = ""
	v.Name = ""
	v.StartTime = ""
	v.EndTime = ""
	v.ActivityType = ""
	v.ModifiedId = ""
	v.ItemId = 0
	v.ActivityId = 0
	v.Status = 0
	v.Version = 0
	poolNrtSceneActivityDto.Put(v)
}
