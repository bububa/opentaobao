package scbp

// AccountEffectDto 结构体
type AccountEffectDto struct {
	// 日期(yyyy-MM-dd)
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 曝光
	Impr string `json:"impr,omitempty" xml:"impr,omitempty"`
	// 点击
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 消耗
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 推广时长
	OnlineMin string `json:"online_min,omitempty" xml:"online_min,omitempty"`
	// 曝光
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 平均点击花费
	ClickCostAvg string `json:"click_cost_avg,omitempty" xml:"click_cost_avg,omitempty"`
	// 单位小时，保留一位小数，例如13.5表示13.5小时
	OnlineTime string `json:"online_time,omitempty" xml:"online_time,omitempty"`
	// 点击量
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
}
