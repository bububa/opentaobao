package axintrade

// PackageRefundPolicyDto 结构体
type PackageRefundPolicyDto struct {
	// 条件退改规则
	ConditionRefundPolicies []ConditionRefundPolicyDto `json:"condition_refund_policies,omitempty" xml:"condition_refund_policies>condition_refund_policy_dto,omitempty"`
	// 退款类型
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}
