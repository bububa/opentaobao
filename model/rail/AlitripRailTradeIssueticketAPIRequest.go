package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriprailtradeissueticketAPIRequest 德铁出票成功接口 API请求
// alitrip.rail.trade.issueticket
//
// 出票成功回调接口
type AlitriprailtradeissueticketAPIRequest struct {
	model.Params
	// 代理商订单号
	_agentOrderId string
	// pnr票号有则填，无则空
	_ticketNo string
	// 平台订单号
	_tpOrderId int64
	// 代理商id
	_agentId int64
}

// NewAlitriprailtradeissueticketRequest 初始化AlitriprailtradeissueticketAPIRequest对象
func NewAlitriprailtradeissueticketRequest() *AlitriprailtradeissueticketAPIRequest {
	return &AlitriprailtradeissueticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriprailtradeissueticketAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.trade.issueticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriprailtradeissueticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriprailtradeissueticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentOrderId is AgentOrderId Setter
// 代理商订单号
func (r *AlitriprailtradeissueticketAPIRequest) SetAgentOrderId(_agentOrderId string) error {
	r._agentOrderId = _agentOrderId
	r.Set("agent_order_id", _agentOrderId)
	return nil
}

// GetAgentOrderId AgentOrderId Getter
func (r AlitriprailtradeissueticketAPIRequest) GetAgentOrderId() string {
	return r._agentOrderId
}

// SetTicketNo is TicketNo Setter
// pnr票号有则填，无则空
func (r *AlitriprailtradeissueticketAPIRequest) SetTicketNo(_ticketNo string) error {
	r._ticketNo = _ticketNo
	r.Set("ticket_no", _ticketNo)
	return nil
}

// GetTicketNo TicketNo Getter
func (r AlitriprailtradeissueticketAPIRequest) GetTicketNo() string {
	return r._ticketNo
}

// SetTpOrderId is TpOrderId Setter
// 平台订单号
func (r *AlitriprailtradeissueticketAPIRequest) SetTpOrderId(_tpOrderId int64) error {
	r._tpOrderId = _tpOrderId
	r.Set("tp_order_id", _tpOrderId)
	return nil
}

// GetTpOrderId TpOrderId Getter
func (r AlitriprailtradeissueticketAPIRequest) GetTpOrderId() int64 {
	return r._tpOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *AlitriprailtradeissueticketAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitriprailtradeissueticketAPIRequest) GetAgentId() int64 {
	return r._agentId
}
