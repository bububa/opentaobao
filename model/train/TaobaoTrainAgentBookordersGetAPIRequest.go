package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentBookordersGetAPIRequest 代理商获取待出票订单列表 API请求
// taobao.train.agent.bookorders.get
//
// 代理商获取待出票订单列表，只返回订单号
type TaobaoTrainAgentBookordersGetAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentBookordersGetRequest 初始化TaobaoTrainAgentBookordersGetAPIRequest对象
func NewTaobaoTrainAgentBookordersGetRequest() *TaobaoTrainAgentBookordersGetAPIRequest {
	return &TaobaoTrainAgentBookordersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookordersGetAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.bookorders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookordersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookordersGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentBookordersGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
