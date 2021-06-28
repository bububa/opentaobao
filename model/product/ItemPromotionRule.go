package product

// ItemPromotionRule 
/* model for simplify = false
type ItemPromotionRule struct {

    // 规则名称
    
    Name   string `json:"name,omitempty"`
    

    // 规则描叙信息
    
    Message   string `json:"message,omitempty"`
    

    // 规则生效开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 规则生效结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 规则类型，常见有SKU锁定规则,下架锁定规则,库存减少锁定规则,库存禁止修改规则,一口价禁止修改规则
    
    Type   string `json:"type,omitempty"`
    

}
*/

// ItemPromotionRule 
type ItemPromotionRule struct {

    // 规则名称
    Name   string `json:"name,omitempty"`

    // 规则描叙信息
    Message   string `json:"message,omitempty"`

    // 规则生效开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 规则生效结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 规则类型，常见有SKU锁定规则,下架锁定规则,库存减少锁定规则,库存禁止修改规则,一口价禁止修改规则
    Type   string `json:"type,omitempty"`

}
