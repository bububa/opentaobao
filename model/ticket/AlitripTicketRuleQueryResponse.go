package ticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票规则信息查询接口 APIResponse
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息
*/
type AlitripTicketRuleQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_ticket_rule_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 门票规则信息
    
    FirstResult   *TicketRuleParam `json:"first_result,omitempty" xml:"