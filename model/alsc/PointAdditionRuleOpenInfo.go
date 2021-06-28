package alsc

// PointAdditionRuleOpenInfo 
/* model for simplify = false
type PointAdditionRuleOpenInfo struct {

    // 可以获取积分的支付方式
    
    AllowPayMethods  struct {
        PayMethodPointAdditionRuleOpenInfo  []PayMethodPointAdditionRuleOpenInfo `json:"pay_method_point_addition_rule_open_info,omitempty"`
    } `json:"allow_pay_methods,omitempty"`
    

    // 不同会员等级的积分获取规则
    
    MemberPointAdditionRules  struct {
        MemberPointAdditionRuleOpenInfo  []MemberPointAdditionRuleOpenInfo `json:"member_point_addition_rule_open_info,omitempty"`
    } `json:"member_point_addition_rules,omitempty"`
    

}
*/

// PointAdditionRuleOpenInfo 
type PointAdditionRuleOpenInfo struct {

    // 可以获取积分的支付方式
    AllowPayMethods   []PayMethodPointAdditionRuleOpenInfo `json:"allow_pay_methods,omitempty"`

    // 不同会员等级的积分获取规则
    MemberPointAdditionRules   []MemberPointAdditionRuleOpenInfo `json:"member_point_addition_rules,omitempty"`

}
