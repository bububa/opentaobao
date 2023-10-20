package xhotel

import (
	"sync"
)

// RateInfo 结构体
type RateInfo struct {
	// 外部价格计划code
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 外部房型id
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
}

var poolRateInfo = sync.Pool{
	New: func() any {
		return new(RateInfo)
	},
}

// GetRateInfo() 从对象池中获取RateInfo
func GetRateInfo() *RateInfo {
	return poolRateInfo.Get().(*RateInfo)
}

// ReleaseRateInfo 释放RateInfo
func ReleaseRateInfo(v *RateInfo) {
	v.RatePlanCode = ""
	v.OutRid = ""
	poolRateInfo.Put(v)
}
