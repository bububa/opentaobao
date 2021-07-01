package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailTradeIssueticketAPIRequest
德铁出票成功接口 API请求
alitrip.rail.trade.issueticket

出票成功回调接口 */
type AlitripRailTradeIssueticketAPIRequest struct {
	model.Params
	// 代理商订单号
	_agentOrderId string
	// 平台订单号
	_tpOrderId int64
	// 代理商id
	_agentId int64
	// pnr票号有则填，无则空
	_ticketNo string
}

// NewAlitripRailTradeIssueticketRequest 初始化AlitripRailTradeIssueticketAPIRequest对象
func NewAlitripRailTradeIssueticketRequest() *AlitripRailTradeIssueticketAPIRequest {
	return &AlitripRailTradeIssueticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailTradeIssueticketAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.trade.issueticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailTradeIssueticketAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AgentOrderId Setter
// 代理商订单号
func (r *AlitripRailTradeIssueticketAPIRequest) SetAgentOrderId(_agentOrderId string) error {
	r._agentOrderId = _agentOrderId
	r.Set("agent_order_id", _agentOrderId)
	return nil
}

// Get AgentOrderId Getter
func (r AlitripRailTradeIssueticketAPIRequest) GetAgentOrderId() string {
	return r._agentOrderId
}

// Set is TpOrderId Setter
// 平台订单号
func (r *AlitripRailTradeIssueticketAPIRequest) SetTpOrderId(_tpOrderId int64) error {
	r._tpOrderId = _tpOrderId
	r.Set("tp_order_id", _tpOrderId)
	return nil
}

// Get TpOrderId Getter
func (r AlitripRailTradeIssueticketAPIRequest) GetTpOrderId() int64 {
	return r._tpOrderId
}

// Set is AgentId Setter
// 代理商id
func (r *AlitripRailTradeIssueticketAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r AlitripRailTradeIssueticketAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// Set is TicketNo Setter
// pnr票号有则填，无则空
func (r *AlitripRailTradeIssueticketAPIRequest) SetTicketNo(_ticketNo string) error {
	r._ticketNo = _ticketNo
	r.Set("ticket_no", _ticketNo)
	return nil
}

// Get TicketNo Getter
func (r AlitripRailTradeIssueticketAPIRequest) GetTicketNo() string {
	return r._ticketNo
}
