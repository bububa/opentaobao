package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeordersGetAPIRequest 获取待改签订单 API请求
// taobao.train.agent.changeorders.get
//
// 代理商用来获取待改签的订单列表及数量，防止代理商掉单。
type TaobaoTrainAgentChangeordersGetAPIRequest struct {
	model.Params
	// 卖家id
	_agentId int64
}

// NewTaobaoTrainAgentChangeordersGetRequest 初始化TaobaoTrainAgentChangeordersGetAPIRequest对象
func NewTaobaoTrainAgentChangeordersGetRequest() *TaobaoTrainAgentChangeordersGetAPIRequest {
	return &TaobaoTrainAgentChangeordersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeordersGetAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.changeorders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeordersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AgentId Setter
// 卖家id
func (r *TaobaoTrainAgentChangeordersGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r TaobaoTrainAgentChangeordersGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
