package simba

// CampaignPlatform 结构体
type CampaignPlatform struct {
	// 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
	SearchChannels []int64 `json:"search_channels,omitempty" xml:"search_channels>int64,omitempty"`
	// 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
	NonsearchChannels []int64 `json:"nonsearch_channels,omitempty" xml:"nonsearch_channels>int64,omitempty"`
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 已经废弃
	OutsideDiscount int64 `json:"outside_discount,omitempty" xml:"outside_discount,omitempty"`
	// 已经废弃了
	MobileDiscount int64 `json:"mobile_discount,omitempty" xml:"mobile_discount,omitempty"`
}
