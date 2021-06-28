package alsc

// MemberPointDeductionRuleOpenInfo 
type MemberPointDeductionRuleOpenInfo struct {

    // 消费金额
    
    ConsumeMoney   int64 `json:"consume_money,omitempty" xml:"consume_money,omitempty"`
    

    // 抵扣积分
    
    DeductedPoint   int64 `json:"deducted_point,omitempty" xml:"deducted_point,omitempty"`
    

    // 是否允许该等级进行积分抵现
    
    Enable   bool `json:"enable,omitempty" xml:"enable,omitempty"`
    

    // 单次使用是否有抵扣上限
    
    HasUpperLimit   bool `json:"has_upper_limit,omitempty" xml:"has_upper_limit,omitempty"`
    

    // 会员等级ID
    
    LevelId   string `json:"level_id,omitempty" xml:"level_id,omitempty"`
    

    // 单次使用的抵扣上限
    
    MaxPoint   int64 `json:"max_point,omitempty" xml:"max_point,omitempty"`
    

}
