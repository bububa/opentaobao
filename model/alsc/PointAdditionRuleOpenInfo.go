package alsc

import (
	"sync"
)

// PointAdditionRuleOpenInfo 结构体
type PointAdditionRuleOpenInfo struct {
	// 可以获取积分的支付方式
	AllowPayMethods []PayMethodPointAdditionRuleOpenInfo `json:"allow_pay_methods,omitempty" xml:"allow_pay_methods>pay_method_point_addition_rule_open_info,omitempty"`
	// 不同会员等级的积分获取规则
	MemberPointAdditionRules []MemberPointAdditionRuleOpenInfo `json:"member_point_addition_rules,omitempty" xml:"member_point_addition_rules>member_point_addition_rule_open_info,omitempty"`
}

var poolPointAdditionRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(PointAdditionRuleOpenInfo)
	},
}

// GetPointAdditionRuleOpenInfo() 从对象池中获取PointAdditionRuleOpenInfo
func GetPointAdditionRuleOpenInfo() *PointAdditionRuleOpenInfo {
	return poolPointAdditionRuleOpenInfo.Get().(*PointAdditionRuleOpenInfo)
}

// ReleasePointAdditionRuleOpenInfo 释放PointAdditionRuleOpenInfo
func ReleasePointAdditionRuleOpenInfo(v *PointAdditionRuleOpenInfo) {
	v.AllowPayMethods = v.AllowPayMethods[:0]
	v.MemberPointAdditionRules = v.MemberPointAdditionRules[:0]
	poolPointAdditionRuleOpenInfo.Put(v)
}
