package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentordergetvtwoAPIRequest 代理商获取订单信息回调APIv2--增加鉴权校验 API请求
// taobao.train.agent.order.get.vtwo
//
// 代理商获取订单信息回调API
type TaobaotrainagentordergetvtwoAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaotrainagentordergetvtwoRequest 初始化TaobaotrainagentordergetvtwoAPIRequest对象
func NewTaobaotrainagentordergetvtwoRequest() *TaobaotrainagentordergetvtwoAPIRequest {
	return &TaobaotrainagentordergetvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentordergetvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentordergetvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentordergetvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaotrainagentordergetvtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagentordergetvtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagentordergetvtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentordergetvtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}
