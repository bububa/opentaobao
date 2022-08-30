package nrt

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
