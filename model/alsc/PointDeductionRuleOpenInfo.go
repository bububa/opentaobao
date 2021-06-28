package alsc

// PointDeductionRuleOpenInfo 
/* model for simplify = false
type PointDeductionRuleOpenInfo struct {

    // 积分扣减规则
    
    MemberPointDeductionRules  struct {
        MemberPointDeductionRuleOpenInfo  []MemberPointDeductionRuleOpenInfo `json:"member_point_deduction_rule_open_info,omitempty"`
    } `json:"member_point_deduction_rules,omitempty"`
    

    // 是否与优惠券同享
    
    WithCoupon   bool `json:"with_coupon,omitempty"`
    

}
*/

// PointDeductionRuleOpenInfo 
type PointDeductionRuleOpenInfo struct {

    // 积分扣减规则
    MemberPointDeductionRules   []MemberPointDeductionRuleOpenInfo `json:"member_point_deduction_rules,omitempty"`

    // 是否与优惠券同享
    WithCoupon   bool `json:"with_coupon,omitempty"`

}
