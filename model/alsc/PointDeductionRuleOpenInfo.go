package alsc

import (
	"sync"
)

// PointDeductionRuleOpenInfo 结构体
type PointDeductionRuleOpenInfo struct {
	// 积分扣减规则
	MemberPointDeductionRules []MemberPointDeductionRuleOpenInfo `json:"member_point_deduction_rules,omitempty" xml:"member_point_deduction_rules>member_point_deduction_rule_open_info,omitempty"`
	// 是否与优惠券同享
	WithCoupon bool `json:"with_coupon,omitempty" xml:"with_coupon,omitempty"`
}

var poolPointDeductionRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(PointDeductionRuleOpenInfo)
	},
}

// GetPointDeductionRuleOpenInfo() 从对象池中获取PointDeductionRuleOpenInfo
func GetPointDeductionRuleOpenInfo() *PointDeductionRuleOpenInfo {
	return poolPointDeductionRuleOpenInfo.Get().(*PointDeductionRuleOpenInfo)
}

// ReleasePointDeductionRuleOpenInfo 释放PointDeductionRuleOpenInfo
func ReleasePointDeductionRuleOpenInfo(v *PointDeductionRuleOpenInfo) {
	v.MemberPointDeductionRules = v.MemberPointDeductionRules[:0]
	v.WithCoupon = false
	poolPointDeductionRuleOpenInfo.Put(v)
}
