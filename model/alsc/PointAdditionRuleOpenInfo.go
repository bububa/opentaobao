package alsc

// PointAdditionRuleOpenInfo 
type PointAdditionRuleOpenInfo struct {
    // 可以获取积分的支付方式
    AllowPayMethods   []PayMethodPointAdditionRuleOpenInfo `json:"allow_pay_methods,omitempty" xml:"allow_pay_methods>pay_method_point_addition_rule_open_info,omitempty"`
    // 不同会员等级的积分获取规则
    MemberPointAdditionRules   []MemberPointAdditionRuleOpenInfo `json:"member_point_addition_rules,omitempty" xml:"member_point_addition_rules>member_point_addition_rule_open_info,omitempty"`
}
