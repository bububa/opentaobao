package xhotelitem

// RatepPlan 结构体
type RatepPlan struct {
	// 系统商
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 房价名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 1：开启2：关闭。
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// ratePlanCode
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
}
