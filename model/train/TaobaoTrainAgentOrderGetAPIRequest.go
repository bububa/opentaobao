package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentordergetAPIRequest 代理商获取订单信息回调API API请求
// taobao.train.agent.order.get
//
// 代理商获取订单信息回调API
type TaobaotrainagentordergetAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaotrainagentordergetRequest 初始化TaobaotrainagentordergetAPIRequest对象
func NewTaobaotrainagentordergetRequest() *TaobaotrainagentordergetAPIRequest {
	return &TaobaotrainagentordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentordergetAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaotrainagentordergetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagentordergetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagentordergetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentordergetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
