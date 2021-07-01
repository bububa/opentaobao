package ieagency

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ieagency"
)

/* 
【国际机票】手工出票 
taobao.alitrip.ie.agent.ticket.issue

代理商手工出票，并回填票号
*/
func TaobaoAlitripIeAgentTicketIssue(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentTicketIssueAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentTicketIssueAPIResponse, error) {
    var resp ieagency.TaobaoAlitripIeAgentTicketIssueAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
