package xhotel

// RateInfo 结构体
type RateInfo struct {
	// 外部价格计划code
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 外部房型id
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
}
