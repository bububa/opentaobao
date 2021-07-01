package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentOrderGetAPIRequest
代理商获取订单信息回调API API请求
taobao.train.agent.order.get

代理商获取订单信息回调API */
type TaobaoTrainAgentOrderGetAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentOrderGetRequest 初始化TaobaoTrainAgentOrderGetAPIRequest对象
func NewTaobaoTrainAgentOrderGetRequest() *TaobaoTrainAgentOrderGetAPIRequest {
	return &TaobaoTrainAgentOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentOrderGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// Get MainOrderId Getter
func (r TaobaoTrainAgentOrderGetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// Set is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentOrderGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r TaobaoTrainAgentOrderGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
