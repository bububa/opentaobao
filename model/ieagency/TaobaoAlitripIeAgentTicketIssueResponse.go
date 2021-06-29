package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】手工出票 APIResponse
taobao.alitrip.ie.agent.ticket.issue

代理商手工出票，并回填票号
*/
type TaobaoAlitripIeAgentTicketIssueAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentTicketIssueResponse
}

type TaobaoAlitripIeAgentTicketIssueResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_ticket_issue_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 回填票号是否成功，true:成功,false：失败
    
    TicketSuccess   bool `json:"ticket_success,omitempty" xml:"ticket_success,omitempty"`

    
}
