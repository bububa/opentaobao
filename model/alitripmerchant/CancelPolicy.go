package alitripmerchant

// CancelPolicy 结构体
type CancelPolicy struct {
	// 取消政策描述
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// 取消政策名称
	CancelPolicyName string `json:"cancel_policy_name,omitempty" xml:"cancel_policy_name,omitempty"`
	// 取消政策类型
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}
