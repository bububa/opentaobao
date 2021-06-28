package ticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票规则信息查询接口 APIResponse
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息
*/
type AlitripTicketRuleQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlitripTicketRuleQueryResponse `json:"alitrip_ticket_rule_query_response,omitempty"` 
    AlitripTicketRuleQueryResponse
}

/* model for simplify = false
type AlitripTicketRuleQueryResponse struct {

    // 门票规则信息
    
    FirstResult  *struct {
        TicketRuleParam  *TicketRuleParam `json:"ticket_rule_param,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type AlitripTicketRuleQueryResponse struct {

    // 门票规则信息
    FirstResult   *TicketRuleParam `json:"first_result,omitempty"`

}
