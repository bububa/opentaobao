package alsc

// PointClearRuleOpenInfo 
/* model for simplify = false
type PointClearRuleOpenInfo struct {

    // 多少天之后清零
    
    Days   int64 `json:"days,omitempty"`
    

    // 积分清零规则类型
    
    PointClearRuleType   string `json:"point_clear_rule_type,omitempty"`
    

}
*/

// PointClearRuleOpenInfo 
type PointClearRuleOpenInfo struct {

    // 多少天之后清零
    Days   int64 `json:"days,omitempty"`

    // 积分清零规则类型
    PointClearRuleType   string `json:"point_clear_rule_type,omitempty"`

}
