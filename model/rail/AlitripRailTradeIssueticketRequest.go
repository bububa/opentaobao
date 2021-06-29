package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
德铁出票成功接口 APIRequest
alitrip.rail.trade.issueticket

出票成功回调接口
*/
type AlitripRailTradeIssueticketRequest struct {
    model.Params

    // 代理商订单号
    agentOrderId   string 

    // 平台订单号
    tpOrderId   int64 

    // 代理商id
    agentId   int64 

    // pnr票号有则填，无则空
    ticketNo   string 

}

func NewAlitripRailTradeIssueticketRequest() *AlitripRailTradeIssueticketRequest{
    return &AlitripRailTradeIssueticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripRailTradeIssueticketRequest) GetApiMethodName() string {
    return "alitrip.rail.trade.issueticket"
}

func (r AlitripRailTradeIssueticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripRailTradeIssueticketRequest) SetAgentOrderId(agentOrderId string) error {
    r.agentOrderId = agentOrderId
    r.Set("agent_order_id", agentOrderId)
    return nil
}

func (r AlitripRailTradeIssueticketRequest) GetAgentOrderId() string {
    return r.agentOrderId
}

func (r *AlitripRailTradeIssueticketRequest) SetTpOrderId(tpOrderId int64) error {
    r.tpOrderId = tpOrderId
    r.Set("tp_order_id", tpOrderId)
    return nil
}

func (r AlitripRailTradeIssueticketRequest) GetTpOrderId() int64 {
    return r.tpOrderId
}

func (r *AlitripRailTradeIssueticketRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlitripRailTradeIssueticketRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *AlitripRailTradeIssueticketRequest) SetTicketNo(ticketNo string) error {
    r.ticketNo = ticketNo
    r.Set("ticket_no", ticketNo)
    return nil
}

func (r AlitripRailTradeIssueticketRequest) GetTicketNo() string {
    return r.ticketNo
}

