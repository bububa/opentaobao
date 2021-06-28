package ticket

// AlitripTicketRuleUploadResultSet 
/* model for simplify = false
type AlitripTicketRuleUploadResultSet struct {

    // 规则维护结果
    
    FirstResult  *struct {
        TopTicketRuleResult  *TopTicketRuleResult `json:"top_ticket_rule_result,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

// AlitripTicketRuleUploadResultSet 
type AlitripTicketRuleUploadResultSet struct {

    // 规则维护结果
    FirstResult   *TopTicketRuleResult `json:"first_result,omitempty"`

}
