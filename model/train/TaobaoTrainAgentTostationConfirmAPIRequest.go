package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagenttostationconfirmAPIRequest 线下票确认送票至车站服务 API请求
// taobao.train.agent.tostation.confirm
//
// 送票至车站的订单，代理商确认配送到站
type TaobaotrainagenttostationconfirmAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaotrainagenttostationconfirmRequest 初始化TaobaotrainagenttostationconfirmAPIRequest对象
func NewTaobaotrainagenttostationconfirmRequest() *TaobaotrainagenttostationconfirmAPIRequest {
	return &TaobaotrainagenttostationconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagenttostationconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.tostation.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagenttostationconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagenttostationconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaotrainagenttostationconfirmAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagenttostationconfirmAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagenttostationconfirmAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagenttostationconfirmAPIRequest) GetAgentId() int64 {
	return r._agentId
}
