package xhotelitem

import (
	"sync"
)

// PromoBindResult 结构体
type PromoBindResult struct {
	// 活动失败原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 外部rid
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
	// 外部rp
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 活动是否报名成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPromoBindResult = sync.Pool{
	New: func() any {
		return new(PromoBindResult)
	},
}

// GetPromoBindResult() 从对象池中获取PromoBindResult
func GetPromoBindResult() *PromoBindResult {
	return poolPromoBindResult.Get().(*PromoBindResult)
}

// ReleasePromoBindResult 释放PromoBindResult
func ReleasePromoBindResult(v *PromoBindResult) {
	v.ErrorMsg = ""
	v.OutRid = ""
	v.RatePlanCode = ""
	v.Success = false
	poolPromoBindResult.Put(v)
}
