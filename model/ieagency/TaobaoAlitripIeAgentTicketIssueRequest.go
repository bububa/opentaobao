package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】手工出票 APIRequest
taobao.alitrip.ie.agent.ticket.issue

代理商手工出票，并回填票号
*/
type TaobaoAlitripIeAgentTicketIssueRequest struct {
    model.Params

    // 代理商id
    agentId   int64 

    // 出票信息
    issueTicketVO   *IeIssueTicketVO 

}

func NewTaobaoAlitripIeAgentTicketIssueRequest() *TaobaoAlitripIeAgentTicketIssueRequest{
    return &TaobaoAlitripIeAgentTicketIssueRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentTicketIssueRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.ticket.issue"
}

func (r TaobaoAlitripIeAgentTicketIssueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentTicketIssueRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentTicketIssueRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoAlitripIeAgentTicketIssueRequest) SetIssueTicketVO(issueTicketVO *IeIssueTicketVO) error {
    r.issueTicketVO = issueTicketVO
    r.Set("issue_ticket_v_o", issueTicketVO)
    return nil
}

func (r TaobaoAlitripIeAgentTicketIssueRequest) GetIssueTicketVO() *IeIssueTicketVO {
    return r.issueTicketVO
}

