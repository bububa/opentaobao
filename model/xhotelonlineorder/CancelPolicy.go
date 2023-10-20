package xhotelonlineorder

import (
	"sync"
)

// CancelPolicy 结构体
type CancelPolicy struct {
	// 取消类型，1:免费取消,2:不可取消,4:从入住时间前推小时前百分比扣款,5:从入住时间前推小时前百分比扣款,6:提前多少小时退款，扣取几晚房费,9:从入住时间前推小时前退订扣款金额
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

var poolCancelPolicy = sync.Pool{
	New: func() any {
		return new(CancelPolicy)
	},
}

// GetCancelPolicy() 从对象池中获取CancelPolicy
func GetCancelPolicy() *CancelPolicy {
	return poolCancelPolicy.Get().(*CancelPolicy)
}

// ReleaseCancelPolicy 释放CancelPolicy
func ReleaseCancelPolicy(v *CancelPolicy) {
	v.CancelPolicyType = 0
	poolCancelPolicy.Put(v)
}
