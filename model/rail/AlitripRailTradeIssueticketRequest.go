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
    _agentOrderId   string
    // 平台订单号
    _tpOrderId   int64
    // 代理商id
    _agentId   int64
    // pnr票号有则填，无则空
    _ticketNo   string
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
func (r *AlitripRailTradeIssueticketRequest) SetAgentOrderId(_agentOrderId string) error {
    r._agentOrderId = _agentOrderId
    r.Set("agent_order_id", _agentOrderId)
    return nil
}

// AgentOrderId Getter
func (r AlitripRailTradeIssueticketRequest) GetAgentOrderId() string {
    return r._agentOrderId
}
// TpOrderId Setter
// 平台订单号
func (r *AlitripRailTradeIssueticketRequest) SetTpOrderId(_tpOrderId int64) error {
    r._tpOrderId = _tpOrderId
    r.Set("tp_order_id", _tpOrderId)
    return nil
}

// TpOrderId Getter
func (r AlitripRailTradeIssueticketRequest) GetTpOrderId() int64 {
    return r._tpOrderId
}
// AgentId Setter
// 代理商id
func (r *AlitripRailTradeIssueticketRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r AlitripRailTradeIssueticketRequest) GetAgentId() int64 {
    return r._agentId
}
// TicketNo Setter
// pnr票号有则填，无则空
func (r *AlitripRailTradeIssueticketRequest) SetTicketNo(_ticketNo string) error {
    r._ticketNo = _ticketNo
    r.Set("ticket_no", _ticketNo)
    return nil
}

// TicketNo Getter
func (r AlitripRailTradeIssueticketRequest) GetTicketNo() string {
    return r._ticketNo
}
