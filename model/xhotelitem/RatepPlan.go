package xhotelitem

import (
	"sync"
)

// RatepPlan 结构体
type RatepPlan struct {
	// 系统商
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 房价名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// ratePlanCode
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 1：开启2：关闭。
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolRatepPlan = sync.Pool{
	New: func() any {
		return new(RatepPlan)
	},
}

// GetRatepPlan() 从对象池中获取RatepPlan
func GetRatepPlan() *RatepPlan {
	return poolRatepPlan.Get().(*RatepPlan)
}

// ReleaseRatepPlan 释放RatepPlan
func ReleaseRatepPlan(v *RatepPlan) {
	v.Vendor = ""
	v.Name = ""
	v.RatePlanCode = ""
	v.Status = 0
	poolRatepPlan.Put(v)
}
