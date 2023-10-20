package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentordergetAPIRequest 【国际机票】查询订单详情 API请求
// taobao.alitrip.ie.agent.order.get
//
// 根据订单ID查询订单详情
type TaobaoalitripieagentordergetAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 交易订单ID
	_tradeOrderId int64
}

// NewTaobaoalitripieagentordergetRequest 初始化TaobaoalitripieagentordergetAPIRequest对象
func NewTaobaoalitripieagentordergetRequest() *TaobaoalitripieagentordergetAPIRequest {
	return &TaobaoalitripieagentordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentordergetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *TaobaoalitripieagentordergetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoalitripieagentordergetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetTradeOrderId is TradeOrderId Setter
// 交易订单ID
func (r *TaobaoalitripieagentordergetAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r TaobaoalitripieagentordergetAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}
