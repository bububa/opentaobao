package scbp

// TopP4pQuickCampaignEffectView 结构体
type TopP4pQuickCampaignEffectView struct {
	// 推广时长
	OnlineTime string `json:"online_time,omitempty" xml:"online_time,omitempty"`
	// 平均点击花费
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 消耗金额，单位元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 点击量
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 曝光量
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}
