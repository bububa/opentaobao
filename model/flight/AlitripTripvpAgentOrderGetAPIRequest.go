package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTripvpAgentOrderGetAPIRequest 廉航辅营正向订单查询详情接口 API请求
// alitrip.tripvp.agent.order.get
//
// 【国际机票】查询辅营订单详情
type AlitripTripvpAgentOrderGetAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 辅营的订单号
	_tradeOrderId int64
}

// NewAlitripTripvpAgentOrderGetRequest 初始化AlitripTripvpAgentOrderGetAPIRequest对象
func NewAlitripTripvpAgentOrderGetRequest() *AlitripTripvpAgentOrderGetAPIRequest {
	return &AlitripTripvpAgentOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderGetAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripTripvpAgentOrderGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetTradeOrderId is TradeOrderId Setter
// 辅营的订单号
func (r *AlitripTripvpAgentOrderGetAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r AlitripTripvpAgentOrderGetAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}
