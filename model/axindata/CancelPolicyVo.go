package axindata

// CancelPolicyVo 结构体
type CancelPolicyVo struct {
	// 详细规则
	PolicyVOList []PolicyVo `json:"policy_v_o_list,omitempty" xml:"policy_v_o_list>policy_vo,omitempty"`
	// 取消政策类型
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}
