package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptripvpagentordergetAPIRequest 廉航辅营正向订单查询详情接口 API请求
// alitrip.tripvp.agent.order.get
//
// 【国际机票】查询辅营订单详情
type AlitriptripvpagentordergetAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 辅营的订单号
	_tradeOrderId int64
}

// NewAlitriptripvpagentordergetRequest 初始化AlitriptripvpagentordergetAPIRequest对象
func NewAlitriptripvpagentordergetRequest() *AlitriptripvpagentordergetAPIRequest {
	return &AlitriptripvpagentordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptripvpagentordergetAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptripvpagentordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptripvpagentordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *AlitriptripvpagentordergetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitriptripvpagentordergetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetTradeOrderId is TradeOrderId Setter
// 辅营的订单号
func (r *AlitriptripvpagentordergetAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r AlitriptripvpagentordergetAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}
