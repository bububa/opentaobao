package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderPayAPIRequest 代购订单代付接口 API请求
// taobao.train.agent.order.pay
//
// 代购订单代付接口
type TaobaoTrainAgentOrderPayAPIRequest struct {
	model.Params
	// 入参对象
	_agentPayOrderParam *AgentPayOrderParam
}

// NewTaobaoTrainAgentOrderPayRequest 初始化TaobaoTrainAgentOrderPayAPIRequest对象
func NewTaobaoTrainAgentOrderPayRequest() *TaobaoTrainAgentOrderPayAPIRequest {
	return &TaobaoTrainAgentOrderPayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderPayAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderPayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentPayOrderParam is AgentPayOrderParam Setter
// 入参对象
func (r *TaobaoTrainAgentOrderPayAPIRequest) SetAgentPayOrderParam(_agentPayOrderParam *AgentPayOrderParam) error {
	r._agentPayOrderParam = _agentPayOrderParam
	r.Set("agent_pay_order_param", _agentPayOrderParam)
	return nil
}

// GetAgentPayOrderParam AgentPayOrderParam Getter
func (r TaobaoTrainAgentOrderPayAPIRequest) GetAgentPayOrderParam() *AgentPayOrderParam {
	return r._agentPayOrderParam
}
