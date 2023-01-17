package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentBookordersGetVtwoAPIRequest 代理商获取待出票订单列表v2--增加鉴权校验 API请求
// taobao.train.agent.bookorders.get.vtwo
//
// 代理商获取待出票订单列表，只返回订单号
type TaobaoTrainAgentBookordersGetVtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentBookordersGetVtwoRequest 初始化TaobaoTrainAgentBookordersGetVtwoAPIRequest对象
func NewTaobaoTrainAgentBookordersGetVtwoRequest() *TaobaoTrainAgentBookordersGetVtwoAPIRequest {
	return &TaobaoTrainAgentBookordersGetVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookordersGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.bookorders.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookordersGetVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentBookordersGetVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookordersGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentBookordersGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}
