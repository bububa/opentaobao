package charity

import (
	"sync"
)

// ChannelUserActionDto 结构体
type ChannelUserActionDto struct {
	// 唯一的动作ID,渠道范围内唯一，幂等控制
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 公益类型子代码
	CharityTypeSubCode string `json:"charity_type_sub_code,omitempty" xml:"charity_type_sub_code,omitempty"`
	// 扩展属性
	FeaturesMap string `json:"features_map,omitempty" xml:"features_map,omitempty"`
	// 公益感言
	Feeling string `json:"feeling,omitempty" xml:"feeling,omitempty"`
	// 做公益当时的时间
	CharityTimestamp string `json:"charity_timestamp,omitempty" xml:"charity_timestamp,omitempty"`
	// 原生公益行为内容，JSON
	OriContent string `json:"ori_content,omitempty" xml:"ori_content,omitempty"`
	// 三方用户标识
	ThirdUserKey string `json:"third_user_key,omitempty" xml:"third_user_key,omitempty"`
	// 版本
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 淘宝ID
	TbUserId int64 `json:"tb_user_id,omitempty" xml:"tb_user_id,omitempty"`
	// 做公益当时的时间戳
	TimestampLong int64 `json:"timestamp_long,omitempty" xml:"timestamp_long,omitempty"`
	// 是否强制报名活动，默认false
	ForeRegActivity bool `json:"fore_reg_activity,omitempty" xml:"fore_reg_activity,omitempty"`
}

var poolChannelUserActionDto = sync.Pool{
	New: func() any {
		return new(ChannelUserActionDto)
	},
}

// GetChannelUserActionDto() 从对象池中获取ChannelUserActionDto
func GetChannelUserActionDto() *ChannelUserActionDto {
	return poolChannelUserActionDto.Get().(*ChannelUserActionDto)
}

// ReleaseChannelUserActionDto 释放ChannelUserActionDto
func ReleaseChannelUserActionDto(v *ChannelUserActionDto) {
	v.EventId = ""
	v.CharityTypeSubCode = ""
	v.FeaturesMap = ""
	v.Feeling = ""
	v.CharityTimestamp = ""
	v.OriContent = ""
	v.ThirdUserKey = ""
	v.Version = ""
	v.ActivityId = 0
	v.TbUserId = 0
	v.TimestampLong = 0
	v.ForeRegActivity = false
	poolChannelUserActionDto.Put(v)
}
