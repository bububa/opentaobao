package alsc

// PointAdditionRuleOpenInfo 
type PointAdditionRuleOpenInfo struct {

    // 可以获取积分的支付方式
    AllowPayMethods   []PayMethodPointAdditionRuleOpenInfo `json:"allow_pay_methods,omitempty"`

    // 不同会员等级的积分获取规则
    MemberPointAdditionRules   []MemberPointAdditionRuleOpenInfo `json:"member_point_addition_rules,omitempty"`

}
