package inventory

// Strategy 
type Strategy struct {

    // 组合货品情况下，是否支持部分子品现货部分子品计划库存可以进行组合。如果不设置，则都是现货库存，或者都是计划库存才能进行组合
    
    CombSupport   bool `json:"comb_support,omitempty" xml:"comb_support,omitempty"`
    

    // 具体的销售策略，1-先现货库存，后计划库存；2-仅计划库存
    
    RuleType   int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
    

}
