package inventory

// StrategyRuleTopDto 
type StrategyRuleTopDto struct {

    // 1，代表先现货后计划库存；2代表仅卖计划库存
    
    RuleType   int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
    

    // 是否支持组合情况下，一部分子品用现货，一部分子品用计划库存。true代表支持，空或者false代表  子品都要同类型的库存才能组合起来
    
    CombSupport   bool `json:"comb_support,omitempty" xml:"comb_support,omitempty"`
    

}
