package wdk

// OpenPromotionRuleStair 
type OpenPromotionRuleStair struct {

    // 阶梯序号，代表商家的阶梯ID，同时代表阶梯优先级的顺序
    Number   int64 `json:"number,omitempty"`

    // 指定件数优惠规则
    CapCountDiscountRule   *CapCountDiscountRule `json:"cap_count_discount_rule,omitempty"`

    // 阶梯整体优惠规则
    CoverAllDiscountRule   *CoverAllDiscountRule `json:"cover_all_discount_rule,omitempty"`

    // 分组表达式【暂时不使用】
    LoginGroupExpress   string `json:"login_group_express,omitempty"`

    // 满x元【单位为分】
    Amount   int64 `json:"amount,omitempty"`

    // 满x件
    Count   int64 `json:"count,omitempty"`

    // 第N件优惠规则
    CountAtDiscountRule   *CountAtDiscountRule `json:"count_at_discount_rule,omitempty"`

    // 是否叠加逻辑分组的条件
    IsOverlayLogicGroupCondition   bool `json:"is_overlay_logic_group_condition,omitempty"`

    // 是否满件
    IsCount   bool `json:"is_count,omitempty"`

    // 是否满元
    IsAmount   bool `json:"is_amount,omitempty"`

    // 单独定价优惠规则
    SeparatePricingDiscountRule   *SeparatePricingDiscountRule `json:"separate_pricing_discount_rule,omitempty"`

}
