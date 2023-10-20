package promotion

import (
	"sync"
)

// CouponApplyResult 结构体
type CouponApplyResult struct {
	// 请求唯一id，问题排查
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 领取结果，领取成功为true，否则为false
	ApplySuccess bool `json:"apply_success,omitempty" xml:"apply_success,omitempty"`
}

var poolCouponApplyResult = sync.Pool{
	New: func() any {
		return new(CouponApplyResult)
	},
}

// GetCouponApplyResult() 从对象池中获取CouponApplyResult
func GetCouponApplyResult() *CouponApplyResult {
	return poolCouponApplyResult.Get().(*CouponApplyResult)
}

// ReleaseCouponApplyResult 释放CouponApplyResult
func ReleaseCouponApplyResult(v *CouponApplyResult) {
	v.TraceId = ""
	v.ApplySuccess = false
	poolCouponApplyResult.Put(v)
}
