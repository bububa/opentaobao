package simba

// Campaign 结构体
type Campaign struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 推广计划名称，不能多余20个汉字
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 推广计划结算状态，offline-下线；online-上线，
	SettleStatus string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 推广计划结算下线原因，1-余额不足；2-超过日限额，以分号分隔多个下线原因
	SettleReason string `json:"settle_reason,omitempty" xml:"settle_reason,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 用户设置的上下限状态；offline-下线；online-上线；
	OnlineStatus string `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 计划类型0 标准计划，16 销量明星
	CampaignType int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
}
