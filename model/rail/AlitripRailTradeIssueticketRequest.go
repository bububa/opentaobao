package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
德铁出票成功接口 API请求
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

// 初始化AlitripRailTradeIssueticketRequest对象
func NewAlitripRailTradeIssueticketRequest() *AlitripRailTradeIssueticketRequest{
    return &AlitripRailTradeIssueticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripRailTradeIssueticketRequest) GetApiMethodName() string {
    return "alitrip.rail.trade.issueticket"
}

// IRequest interface 方法, 获取API参数
func (r AlitripRailTradeIssueticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentOrderId Setter
// 代理商订单号
func (r *AlitripRailTradeIssueticketRequest) SetAgentOrderId(agentOrderId string) error {
    r.agentOrderId = agentOrderId
    r.Set("agent_order_id", agentOrderId)
    return nil
}

// AgentOrderId Getter
func (r AlitripRailTradeIssueticketRequest) GetAgentOrderId() string {
    return r.agentOrderId
}
// TpOrderId Setter
// 平台订单号
func (r *AlitripRailTradeIssueticketRequest) SetTpOrderId(tpOrderId int64) error {
    r.tpOrderId = tpOrderId
    r.Set("tp_order_id", tpOrderId)
    return nil
}

// TpOrderId Getter
func (r AlitripRailTradeIssueticketRequest) GetTpOrderId() int64 {
    return r.tpOrderId
}
// AgentId Setter
// 代理商id
func (r *AlitripRailTradeIssueticketRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r AlitripRailTradeIssueticketRequest) GetAgentId() int64 {
    return r.agentId
}
// TicketNo Setter
// pnr票号有则填，无则空
func (r *AlitripRailTradeIssueticketRequest) SetTicketNo(ticketNo string) error {
    r.ticketNo = ticketNo
    r.Set("ticket_no", ticketNo)
    return nil
}

// TicketNo Getter
func (r AlitripRailTradeIssueticketRequest) GetTicketNo() string {
    return r.ticketNo
}
