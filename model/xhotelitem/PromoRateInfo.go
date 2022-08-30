package xhotelitem

// PromoRateInfo 结构体
type PromoRateInfo struct {
	// 外部rp
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 外部rid
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
}
