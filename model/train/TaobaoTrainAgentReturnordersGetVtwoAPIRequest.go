package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentreturnordersgetvtwoAPIRequest 获取待退票的订单v2--增加鉴权校验 API请求
// taobao.train.agent.returnorders.get.vtwo
//
// 代理商用来获取待退票的订单列表及数量，防止代理商掉单。
type TaobaotrainagentreturnordersgetvtwoAPIRequest struct {
	model.Params
	// 卖家ID
	_agentId int64
	// 0 线上退票 1线下退票
	_offline int64
}

// NewTaobaotrainagentreturnordersgetvtwoRequest 初始化TaobaotrainagentreturnordersgetvtwoAPIRequest对象
func NewTaobaotrainagentreturnordersgetvtwoRequest() *TaobaotrainagentreturnordersgetvtwoAPIRequest {
	return &TaobaotrainagentreturnordersgetvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentreturnordersgetvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnorders.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentreturnordersgetvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentreturnordersgetvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 卖家ID
func (r *TaobaotrainagentreturnordersgetvtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentreturnordersgetvtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetOffline is Offline Setter
// 0 线上退票 1线下退票
func (r *TaobaotrainagentreturnordersgetvtwoAPIRequest) SetOffline(_offline int64) error {
	r._offline = _offline
	r.Set("offline", _offline)
	return nil
}

// GetOffline Offline Getter
func (r TaobaotrainagentreturnordersgetvtwoAPIRequest) GetOffline() int64 {
	return r._offline
}
