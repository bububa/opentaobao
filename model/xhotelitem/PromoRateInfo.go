package xhotelitem

import (
	"sync"
)

// PromoRateInfo 结构体
type PromoRateInfo struct {
	// 外部rp
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 外部rid
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
}

var poolPromoRateInfo = sync.Pool{
	New: func() any {
		return new(PromoRateInfo)
	},
}

// GetPromoRateInfo() 从对象池中获取PromoRateInfo
func GetPromoRateInfo() *PromoRateInfo {
	return poolPromoRateInfo.Get().(*PromoRateInfo)
}

// ReleasePromoRateInfo 释放PromoRateInfo
func ReleasePromoRateInfo(v *PromoRateInfo) {
	v.RatePlanCode = ""
	v.OutRid = ""
	poolPromoRateInfo.Put(v)
}
