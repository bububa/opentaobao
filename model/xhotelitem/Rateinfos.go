package xhotelitem

// Rateinfos 结构体
type Rateinfos struct {
	// 外部房源id
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
	// 外部rpcode
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
}
