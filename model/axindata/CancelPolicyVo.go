package axindata

import (
	"sync"
)

// CancelPolicyVo 结构体
type CancelPolicyVo struct {
	// 详细规则
	PolicyVOList []PolicyVo `json:"policy_v_o_list,omitempty" xml:"policy_v_o_list>policy_vo,omitempty"`
	// 取消政策类型
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

var poolCancelPolicyVo = sync.Pool{
	New: func() any {
		return new(CancelPolicyVo)
	},
}

// GetCancelPolicyVo() 从对象池中获取CancelPolicyVo
func GetCancelPolicyVo() *CancelPolicyVo {
	return poolCancelPolicyVo.Get().(*CancelPolicyVo)
}

// ReleaseCancelPolicyVo 释放CancelPolicyVo
func ReleaseCancelPolicyVo(v *CancelPolicyVo) {
	v.PolicyVOList = v.PolicyVOList[:0]
	v.CancelPolicyType = 0
	poolCancelPolicyVo.Put(v)
}
