package travel

// BookingRuleInfo 
/* model for simplify = false
type BookingRuleInfo struct {

    // 1500个字
    
    RuleDesc   string `json:"rule_desc,omitempty"`
    

    // fee_included：费用包含，跟团游必填； fee_excluded：费用不含，所有类目必填； order_info：预定须知； extra_cost：其他费用，预留；
    
    RuleType   string `json:"rule_type,omitempty"`
    

    // 规则说明，1500个字符
    
    RuleList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"rule_list,omitempty"`
    

}
*/

// BookingRuleInfo 
type BookingRuleInfo struct {

    // 1500个字
    RuleDesc   string `json:"rule_desc,omitempty"`

    // fee_included：费用包含，跟团游必填； fee_excluded：费用不含，所有类目必填； order_info：预定须知； extra_cost：其他费用，预留；
    RuleType   string `json:"rule_type,omitempty"`

    // 规则说明，1500个字符
    RuleList   []string `json:"rule_list,omitempty"`

}
