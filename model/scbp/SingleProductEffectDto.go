package scbp

// SingleProductEffectDto 结构体
type SingleProductEffectDto struct {
	// 产品title
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 曝光
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 日期
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 平均点击花费
	ClickCostAvg string `json:"click_cost_avg,omitempty" xml:"click_cost_avg,omitempty"`
	// 点击数
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 单位元，保留两位小数, 例如3.75表示3.75元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 百分比，保留两位小数，例如3.75表示3.75%
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
