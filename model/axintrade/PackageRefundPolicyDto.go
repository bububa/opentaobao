package axintrade

import (
	"sync"
)

// PackageRefundPolicyDto 结构体
type PackageRefundPolicyDto struct {
	// 条件退改规则
	ConditionRefundPolicies []ConditionRefundPolicyDto `json:"condition_refund_policies,omitempty" xml:"condition_refund_policies>condition_refund_policy_dto,omitempty"`
	// 退款类型
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

var poolPackageRefundPolicyDto = sync.Pool{
	New: func() any {
		return new(PackageRefundPolicyDto)
	},
}

// GetPackageRefundPolicyDto() 从对象池中获取PackageRefundPolicyDto
func GetPackageRefundPolicyDto() *PackageRefundPolicyDto {
	return poolPackageRefundPolicyDto.Get().(*PackageRefundPolicyDto)
}

// ReleasePackageRefundPolicyDto 释放PackageRefundPolicyDto
func ReleasePackageRefundPolicyDto(v *PackageRefundPolicyDto) {
	v.ConditionRefundPolicies = v.ConditionRefundPolicies[:0]
	v.RefundType = 0
	poolPackageRefundPolicyDto.Put(v)
}
