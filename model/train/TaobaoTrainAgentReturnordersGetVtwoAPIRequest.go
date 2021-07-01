package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentReturnordersGetVtwoAPIRequest
获取待退票的订单v2--增加鉴权校验 API请求
taobao.train.agent.returnorders.get.vtwo

代理商用来获取待退票的订单列表及数量，防止代理商掉单。 */
type TaobaoTrainAgentReturnordersGetVtwoAPIRequest struct {
	model.Params
	// 卖家ID
	_agentId int64
	// 0 线上退票 1线下退票
	_offline int64
}

// NewTaobaoTrainAgentReturnordersGetVtwoRequest 初始化TaobaoTrainAgentReturnordersGetVtwoAPIRequest对象
func NewTaobaoTrainAgentReturnordersGetVtwoRequest() *TaobaoTrainAgentReturnordersGetVtwoAPIRequest {
	return &TaobaoTrainAgentReturnordersGetVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnorders.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AgentId Setter
// 卖家ID
func (r *TaobaoTrainAgentReturnordersGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// Set is Offline Setter
// 0 线上退票 1线下退票
func (r *TaobaoTrainAgentReturnordersGetVtwoAPIRequest) SetOffline(_offline int64) error {
	r._offline = _offline
	r.Set("offline", _offline)
	return nil
}

// Get Offline Getter
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetOffline() int64 {
	return r._offline
}
