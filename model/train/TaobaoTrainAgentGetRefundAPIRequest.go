package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentgetrefundAPIRequest 代理商获取订单退票信息 API请求
// taobao.train.agent.get.refund
//
// 代理商获取订单信息回调API
type TaobaotrainagentgetrefundAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaotrainagentgetrefundRequest 初始化TaobaotrainagentgetrefundAPIRequest对象
func NewTaobaotrainagentgetrefundRequest() *TaobaotrainagentgetrefundAPIRequest {
	return &TaobaotrainagentgetrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentgetrefundAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.get.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentgetrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentgetrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaotrainagentgetrefundAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagentgetrefundAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagentgetrefundAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentgetrefundAPIRequest) GetAgentId() int64 {
	return r._agentId
}
