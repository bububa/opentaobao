package alsc

// PointDeductionRuleOpenInfo 
type PointDeductionRuleOpenInfo struct {

    // 积分扣减规则
    MemberPointDeductionRules   []MemberPointDeductionRuleOpenInfo `json:"member_point_deduction_rules,omitempty"`

    // 是否与优惠券同享
    WithCoupon   bool `json:"with_coupon,omitempty"`

}
