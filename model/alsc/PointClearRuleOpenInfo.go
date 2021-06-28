package alsc

// PointClearRuleOpenInfo 
type PointClearRuleOpenInfo struct {

    // 多少天之后清零
    
    Days   int64 `json:"days,omitempty" xml:"days,omitempty"`
    

    // 积分清零规则类型
    
    PointClearRuleType   string `json:"point_clear_rule_type,omitempty" xml:"point_clear_rule_type,omitempty"`
    

}
